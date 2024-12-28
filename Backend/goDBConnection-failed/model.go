package main

import (
	"gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `gorm:"size:100;not null"`
}

// MigrateUser migrates the User model.
func MigrateUser(db *gorm.DB) {
    db.AutoMigrate(&User{})
}

