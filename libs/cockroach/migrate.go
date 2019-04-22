package cockroach

import (
	"github.com/jinzhu/gorm"
)

// Migrate all models.
func migrate(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}
