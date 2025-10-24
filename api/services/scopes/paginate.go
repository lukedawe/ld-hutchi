package scopes

import (
	"gorm.io/gorm"
)

func Paginate(page uint, pageSize uint) func(db *gorm.Statement) {
	return func(db *gorm.Statement) {
		if page == 0 {
			page = 1
		}

		if pageSize > 100 {
			pageSize = 100
		}

		offset := (page - 1) * pageSize
		db.Offset(int(offset)).Limit(int(pageSize))
	}
}
