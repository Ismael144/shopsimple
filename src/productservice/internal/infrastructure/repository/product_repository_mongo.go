package repository

import (
	"context"
	"math"
	"time"

	"github.com/Ismael144/productservice/internal/application/ports"
	"github.com/Ismael144/productservice/internal/domain"
	entities "github.com/Ismael144/productservice/internal/domain/entities"
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"github.com/Ismael144/productservice/internal/infrastructure/repository/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: db.Collection("products"),
	}
}

// Create product
func (r *MongoProductRepository) Create(
	ctx context.Context,
	product *entities.Product,
) error {

	doc := ProductDomainToModel(product)

	_, err := r.collection.InsertOne(ctx, doc)
	return err
}

// List products with pagination
func (r *MongoProductRepository) List(
	ctx context.Context,
	page, pageSize uint32,
) ([]*entities.Product, *entities.Pagination, error) {

	skip := int64((page - 1) * pageSize)
	limit := int64(pageSize)

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, nil, err
	}

	var docs []product.ProductModelMongo
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, nil, err
	}

	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}

	// Get total pages
	totalPages := math.Ceil(float64(total / limit))

	products := make([]*entities.Product, 0, len(docs))
	for _, d := range docs {
		products = append(products, ProductModelToDomain(&d))
	}

	return products, &entities.Pagination{
		CurrentPage: int64(page),
		TotalPages:  int64(totalPages),
		TotalItems:  int64(total),
	}, nil
}

// Update stock
func (r *MongoProductRepository) UpdateStock(
	ctx context.Context,
	productID *valueobjects.ProductID,
	stock int64,
) error {
	res, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": productID.String()},
		bson.M{
			"$set": bson.M{
				"stock":      stock,
				"updated_at": time.Now(),
			},
		},
	)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return domain.ErrProductNotFound
	}

	return nil
}

// Find product by id from db
func (r *MongoProductRepository) FindById(
	ctx context.Context,
	productID *valueobjects.ProductID,
) (*entities.Product, error) {

	var doc product.ProductModelMongo

	err := r.collection.FindOne(
		ctx,
		bson.M{"_id": productID.String()},
	).Decode(&doc)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrProductNotFound
		}
		return nil, err
	}

	return ProductModelToDomain(&doc), nil
}

// TODO: Implement Filter
func (r *MongoProductRepository) Filter(
	ctx context.Context,
	filters *ports.ProductFilters,
) ([]*entities.Product, *entities.Pagination, error) {

	query := bson.M{}

	// Apply search string filter - search in both name and description
	if filters.SearchString != "" {
		searchRegex := primitive.Regex{
			Pattern: filters.SearchString,
			Options: "i", // Case-insensitive search
		}
		query["$or"] = bson.A{
			bson.M{"name": searchRegex},
			bson.M{"description": searchRegex},
		}
	}

	// Apply categories filter
	if len(filters.Categories) > 0 {
		query["categories"] = bson.M{
			"$in": filters.Categories,
		}
	}

	// Apply price changes filter
	// We check whether both min and max ranges are not zero
	// if so, then we proceed to apply filter
	if filters.PriceRanges != nil && (!valueobjects.IsZero(filters.PriceRanges.Min) && !valueobjects.IsZero(filters.PriceRanges.Max)) {
		query["price_minor"] = bson.M{
			"$gte": moneyToMinor(filters.PriceRanges.Min),
			"$lte": moneyToMinor(filters.PriceRanges.Max),
		}
	}

	skip := int64((filters.Page - 1) * filters.PageSize)
	limit := int64(filters.PageSize)

	opts := options.Find().
		SetSkip(skip).
		SetLimit(limit).
		SetSort(bson.M{"created_at": -1})

	cursor, err := r.collection.Find(ctx, query, opts)
	if err != nil {
		return nil, nil, err
	}

	var docs []product.ProductModelMongo
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, nil, err
	}

	total, err := r.collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, nil, err
	}

	// Get total pages
	totalPages := math.Ceil(float64(total / limit))

	products := make([]*entities.Product, 0, len(docs))
	for _, d := range docs {
		products = append(products, ProductModelToDomain(&d))
	}

	return products, &entities.Pagination{
		CurrentPage: int64(filters.Page),
		TotalPages:  int64(totalPages),
		TotalItems:  total,
	}, nil
}

func (r *MongoProductRepository) BatchFindById(
	ctx context.Context,
	productIDs []*valueobjects.ProductID,
) ([]*entities.Product, int64, error) {

	objIDs := make([]string, 0, len(productIDs))
	for _, id := range productIDs {
		objIDs = append(objIDs, id.String())
	}

	cursor, err := r.collection.Find(
		ctx,
		bson.M{"_id": bson.M{"$in": objIDs}},
	)
	if err != nil {
		return nil, 0, err
	}

	var docs []product.ProductModelMongo
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, 0, err
	}

	products := make([]*entities.Product, 0, len(docs))
	for _, d := range docs {
		products = append(products, ProductModelToDomain(&d))
	}

	return products, int64(len(products)), nil
}

func (r *MongoProductRepository) ListCategories(
	ctx context.Context,
) ([]string, error) {
	// Init array of categories
	var categories []string

	err := r.collection.Distinct(ctx, "categories", bson.D{}).Decode(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
