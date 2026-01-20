package repository

import (
	"context"
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

// Helper function to run operation on cart
// Saving your from loading and storing cart json each time you mutate the cart
func mapCart(ctx context.Context, rdb *redis.Client, userID valueobjects.UserID, op func(cart *entities.Cart) (*entities.Cart, error)) (*entities.Cart, error) {
	exists, err := rdb.Exists(ctx, userID.Key()).Result()
	if err != nil {
		return nil, err
	}
	if exists == 1 {
		cartJson, err := rdb.Get(ctx, userID.Key()).Result()
		if err != nil {
			return nil, err
		}
		// Unmarshal cart json
		cart, err := entities.UnmarshalCart(cartJson)
		if err != nil {
			return nil, err
		}
		cart, err = op(cart)
		if err != nil {
			return nil, err
		}
		// We marshal the cart again for storage
		cartJson, err = cart.Marshal()
		if err != nil {
			return nil, err
		}
		// Marshal and then store with a ttl of a day.
		_, err = rdb.Set(ctx, userID.Key(), []byte(cartJson), time.Duration(24*time.Hour)).Result()
		if err != nil {
			return nil, err
		}

		return cart, nil
	} else {
		// Else create a new cart for user
		cart := entities.NewCart(userID)
		cart, err = op(cart)
		cartJson, err := cart.Marshal()
		if err != nil {
			return nil, err
		}
		// Store with ttl of a day
		if _, err = rdb.Set(ctx, userID.Key(), []byte(cartJson), time.Duration(24*time.Hour)).Result(); err != nil {
			return nil, err
		}
		return cart, nil
	}
}

// Check whether an item is already in cart
func (r *CartRepositoryRedis) IsItemInCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.CartItem, error) {
	var cartItem *entities.CartItem

	_, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		cartItem = cart.GetById(productID)
		return cart, nil
	})
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// Add item into cart
// It takes in productStock value which will
// be provided by the productservice, returns
// ErrStockLimitReached when quantity exceeds
// the given product's stock
func (r *CartRepositoryRedis) AddItem(ctx context.Context, userID valueobjects.UserID, item *entities.CartItem, productStock uint32) (*entities.Cart, error) {
	cart, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		cart, err := cart.AddToCart(item, productStock)
		if err != nil {
			return nil, err
		}
		return cart, err
	})
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// Deduct item quantity in cart
func (r *CartRepositoryRedis) DeductFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID, Quantity uint32) (*entities.Cart, error) {
	// Call mapCart
	cart, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		cart = cart.DeductFromCart(productID, Quantity)
		return cart, nil
	})
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// Remove item from cart
func (r *CartRepositoryRedis) RemoveFromCart(ctx context.Context, userID valueobjects.UserID, productID valueobjects.ProductID) (*entities.Cart, error) {
	cart, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		cart = cart.RemoveItem(productID)
		return cart, nil
	})
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// If user authenticates, guest id which is generated from frontend to as key for cart
// Is replaced by id of authenticated user
func (r *CartRepositoryRedis) AssignToAuthUser(ctx context.Context, guestUserID valueobjects.UserID, authUserID valueobjects.UserID) error {
	// Check for user id existence
	exists, err := r.rdb.Exists(ctx, guestUserID.Key()).Result()
	if err != nil {
		return err
	}
	if exists == 1 {
		// Get cart
		cart, err := r.GetCart(ctx, guestUserID)
		if err != nil {
			return err
		}
		// Marshal cart
		cartJson, err := cart.Marshal()
		if err != nil {
			return err
		}
		// We assign guest cart to auth user
		if _, err = r.rdb.Set(ctx, authUserID.Key(), []byte(cartJson), time.Duration(24*time.Hour)).Result(); err != nil {
			return err
		}
	} else {
		// If no guest user, then we create new cart and assign it authuser id
		cart := entities.NewCart(authUserID)
		cartJson, err := cart.Marshal()
		if err != nil {
			return err
		}
		if _, err = r.rdb.Set(ctx, authUserID.Key(), []byte(cartJson), time.Duration(24*time.Hour)).Result(); err != nil {
			return err
		}
	}
	return nil
}

// Clear all items in cart
func (r *CartRepositoryRedis) EmptyCart(ctx context.Context, userID valueobjects.UserID) error {
	_, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		cart = cart.Empty()
		return cart, nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Return current state of the cart with actual cart price total
func (r *CartRepositoryRedis) GetCart(ctx context.Context, userID valueobjects.UserID) (*entities.Cart, error) {
	// Init cart
	var cart *entities.Cart
	cart, err := mapCart(ctx, r.rdb, userID, func(cart *entities.Cart) (*entities.Cart, error) {
		return cart.GetCart(), nil
	})
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// Save current state of current by user id in our database... 
func (r *CartRepositoryRedis) Save(ctx context.Context, userID valueobjects.UserID, cart *entities.Cart) error {
	_, err := mapCart(ctx, r.rdb, userID, func(_ *entities.Cart) (*entities.Cart, error) {
		return cart, nil
	})
	if err != nil {
		return err
	}
	return nil
}
