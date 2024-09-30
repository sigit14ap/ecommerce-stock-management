package dto

import "time"

type ProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
}

type ProductResponse struct {
	ID         uint64    `json:"id"`
	ShopID     uint64    `json:"shop_id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	TotalStock int64     `json:"total_stock"`
}
