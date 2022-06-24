package models

type Order struct {
	Id         uint   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
	OrderItems []OrderItem
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float64 `json:"price"`
	Quantity     uint    `json:"quantity"`
}
