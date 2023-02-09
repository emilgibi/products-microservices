package models

type Product struct {
	ID              int    `json:"id"`
	ProductName     string `json:"product_name"`
	ProductQuantity string `json:"product_quantity"`
}
