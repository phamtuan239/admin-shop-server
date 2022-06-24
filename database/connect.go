package database

import (
	"admin.server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "postgresql://postgres:phamtuan239@localhost:5432/go_admin"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connect to postgres: " + err.Error())
	}

	DB = db

	// db.Migrator().DropTable(&models.Role{})
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable(&models.Permission{})
	// db.Migrator().DropTable(&models.Permission{})
	// db.Migrator().DropTable(&models.Product{})

	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
