package cockroach

import (
	"go-microservices/libs/logger"
	"time"
)

type Post struct {
	ID        int64 `gorm:"primary_key"`
	Title     string
	CreatedAt time.Time
}

func GetAllPosts() ([]Post, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	err = db.Find(&posts).Error
	if err != nil {
		logger.GetCockroach().Error(11, map[string]interface{}{
			"error":  err,
			"method": "GetAllPosts",
		})
		return nil, err
	}

	return posts, nil
}

func GetPost(id int64) (*Post, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	post := Post{}
	err = db.Where(&Post{ID: id}).First(&post).Error
	if err != nil {
		logger.GetCockroach().Error(11, map[string]interface{}{
			"error":  err,
			"method": "GetPost",
			"id":     id,
		})
		return nil, err
	}

	return &post, nil
}

func CreatePost(p Post) (*Post, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}

	post := Post{}
	err = db.Create(&Post{
		CreatedAt: time.Now(),
		Title:     p.Title,
	}).Scan(&post).Error
	if err != nil {
		logger.GetCockroach().Error(11, map[string]interface{}{
			"error":  err,
			"method": "CreatePost",
			"post":   p,
		})
		return nil, err
	}

	return &post, nil
}

func UpdatePost(q Post, u Post) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	err = db.Model(&q).Updates(&Post{
		Title:     u.Title,
		CreatedAt: time.Now(),
	}).Error
	if err != nil {
		logger.GetCockroach().Error(11, map[string]interface{}{
			"error":  err,
			"method": "UpdatePost",
			"query":  q,
			"update": u,
		})
		return err
	}

	return nil
}

func DeletePost(id int64) error {
	db, err := Connection()
	if err != nil {
		return err
	}

	err = db.Delete(&Post{ID: id}).Error
	if err != nil {
		logger.GetCockroach().Error(11, map[string]interface{}{
			"error":  err,
			"method": "DeletePost",
			"id":     id,
		})
		return err
	}

	return nil
}
