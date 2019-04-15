package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import "fmt"

func GetPost(id int64) (*Post, error) {
	var post Post
	result := DB.First(&post, id)

	if result.Error != nil {
		return nil, fmt.Errorf("Post with id '%d' was NOT found.", id)
	}

	return &post, nil
}

func GetPosts() ([]Post, error) {
	var posts []Post
	result := DB.Find(&posts)

	return posts, result.Error
}

func CreatePost(post Post) (Post, error) {
	result := DB.Create(&post)

	return post, result.Error
}

func DeletePostById(id int64) error {
	result := DB.Where("id = ?", id).Delete(&Post{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("Post with id '%d' was NOT found.", id)
	}

	return result.Error
}
