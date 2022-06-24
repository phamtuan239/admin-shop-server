package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

func (user *User) SetPassword() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashPassword)
}

// db.Preload("Role").Offset(offset).Limit(limit).Find(&users)

func (user *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Find(&User{}).Count(&total)
	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	users := []User{}

	db.Preload("Role").Preload("Permissions").Offset(offset).Limit(limit).Find(&users)

	return users
}
