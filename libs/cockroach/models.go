package cockroach

import "github.com/jinzhu/gorm"

type Post struct {
	Title string
	Date  string
	db    *gorm.DB
}
