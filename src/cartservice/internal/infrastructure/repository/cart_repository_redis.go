package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Ismael144/cartservice/internal/domain/entities"
	"github.com/Ismael144/cartservice/internal/domain/valueobjects"
	"github.com/redis/go-redis/v9"
)

type CartRepositoryRedis struct {
	rdb *redis.Client
}

func NewCartRepositoryRedis(rdb *redis.Client) *CartRepositoryRedis {
	return &CartRepositoryRedis{rdb: rdb}
}

func (r *CartRepositoryRedis) AddItem(ctx context.Context, userID valueobjects.UserID, item *entities.CartItem) (*valueobjects.ProductID, error) {
	exists, err := r.rdb.Exists(ctx, userID.Key()).Result()
	if err != nil {
		return nil, err
	}

	if exists == 1 {
		cartJson, err := r.rdb.Get(ctx, userID.Key()).Result()
		fmt.Println("Cart Json:", cartJson)
		if err != nil {
			return nil, err
		}
		// Unmarshal cart json
		cart, err := entities.UnmarshalCart(cartJson)
		if err != nil {
			return nil, err
		}
		cart.AddToCart(item)
	} else {
		cart := entities.NewCart(valueobjects.ParseUserID(userID.String()))
		// Add item to cart
		cart.AddToCart(item)
		// Jsonify the cart
		cartJson, err := cart.Marshal()
		if err != nil {
			return nil, err
		}

		if _, err = r.rdb.Set(ctx, userID.Key(), []byte(cartJson), time.Duration(24*time.Hour)).Result(); err != nil {
			return nil, err
		}
	}



	return &item.ProductID, nil
}

// func RemoveFrom