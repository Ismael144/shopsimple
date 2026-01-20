package repository

import (
	"github.com/Ismael144/productservice/internal/domain/valueobjects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func productIDToObjectID(id valueobjects.ProductID) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}