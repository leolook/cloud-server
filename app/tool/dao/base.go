package dao

import "github.com/jinzhu/gorm"

type Base struct {
	DB *gorm.DB
}

func (b *Base) checkRow(row *gorm.DB) (bool, error) {

	if row.RecordNotFound() {
		return false, nil
	}

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}
