package util

import (
	"gorm.io/gorm"
)

func Paginate(page uint, pageSize uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		if pageSize > 100 {
			pageSize = 100
		}

		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
