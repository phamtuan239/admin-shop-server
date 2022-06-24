package models

import (
	"admin.server/helpers"
	"gorm.io/gorm"
)

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (product *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Find(&Product{}).Count(&total)
	return total
}

func (product *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	products := []Product{}

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}

func (product *Product) UploadFile(input interface{}) error {

	image, err := helpers.UploadFileToCloudinary(input, "products")

	if err != nil {
		return err
	}

	product.Image = image

	return nil
}
