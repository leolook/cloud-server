package dao

import "github.com/jinzhu/gorm"

type Base struct {
	DB *gorm.DB
}
