package models

import (
	"errors"
	"time"
)

// Product datastructure for Poduct domain
type Product struct {
	id          string
	name        string
	price       int
	brand       Brand
	description string
	stock       int
	sold        int
	createdAt   time.Time
	updatedAt   time.Time
}

// BuyStock will subtract product stock and add product sold
func (product *Product) BuyStock(amount int) (int, error) {
	if product.stock < amount {
		return 0, errors.New("Insufficient stock")
	}

	product.stock = product.stock - amount
	product.sold = product.sold + amount
	product.updatedAt = time.Now()

	return product.stock, nil
}

// AddStock will add product stock
func (product *Product) AddStock(amount int) (int, error) {
	product.stock = product.stock + amount
	product.updatedAt = time.Now()

	return product.stock, nil
}
