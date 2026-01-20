package domain

import "errors"

var (
	ErrInvalidQuantity = errors.New("Invalid quantity")
	ErrStockLimitReached = errors.New("stock limit reached")
)