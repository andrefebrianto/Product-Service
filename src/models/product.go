package models

import (
	"errors"
	"time"
)

// Product datastructure for Poduct domain
type Product struct {
	ID          string
	Name        string
	Price       int
	Brand       Brand
	Description string
	Stock       int
	Sold        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// BuyStock will subtract product stock and add product sold
func (product *Product) BuyStock(amount int) (int, error) {
	if product.Stock < amount {
		return 0, errors.New("Insufficient stock")
	}

	product.Stock = product.Stock - amount
	product.Sold = product.Sold + amount
	product.UpdatedAt = time.Now()

	return product.Stock, nil
}

// AddStock will add product stock
func (product *Product) AddStock(amount int) (int, error) {
	product.Stock = product.Stock + amount
	product.UpdatedAt = time.Now()

	return product.Stock, nil
}
