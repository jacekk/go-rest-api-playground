package database

import (
	"fmt"
)

func GetPost(id uint64) (*Post, error) {
	var post Post
	result := DB.Preload("Category").First(&post, id)

	if result.Error != nil {
		return nil, fmt.Errorf("Post with id '%d' was NOT found.", id)
	}

	return &post, nil
}

func GetPosts(offset uint64, limit uint64) ([]Post, error) {
	var posts []Post
	// NOTE: offset and limit has to be before find
	result := DB.Preload("Category").Offset(offset).Limit(limit).Find(&posts)

	return posts, result.Error
}

func GetAuthorPosts(offset uint64, limit uint64, userId uint64) ([]Post, error) {
	var posts []Post
	where := &Post{AuthorID: uint(userId)}
	result := DB.Preload("Category").Offset(offset).Limit(limit).Where(where).Find(&posts)

	return posts, result.Error
}

func CreatePost(post Post) (Post, error) {
	result := DB.Create(&post)

	return post, result.Error
}

func DeletePostById(id uint64) error {
	result := DB.Where("id = ?", id).Delete(&Post{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("Post with id '%d' was NOT found.", id)
	}

	return result.Error
}
