package models

import (
	"errors"
	"time"
)

// Product datastructure for Poduct domain
type Product struct {
	ID          string    `pg:"id,type:uuid,pk"`
	Name        string    `pg:"name,notnull"`
	Price       int       `pg:"price,notnull"`
	Brand       *Brand    `pg:"rel:has-one,notnull,join_fk:brand_id"`
	Description string    `pg:"description"`
	Stock       int       `pg:"stock,notnull"`
	Sold        int       `pg:"sold,notnull"`
	CreatedAt   time.Time `pg:"created_at,notnull,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at,notnull,default:now()"`
	DeletedAt   time.Time `pg:"deleted_at,soft_delete"`
}

// BuyProduct will subtract product stock and add product sold
func (product *Product) BuyProduct(amount int) (int, error) {
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
