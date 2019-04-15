package database // import "github.com/jacekk/go-rest-api-playground/internal/database

func GetPosts() ([]Post, error) {
	var posts []Post
	result := DB.Find(&posts)

	return posts, result.Error
}

func CreatePost(post Post) (Post, error) {
	result := DB.Create(&post)

	return post, result.Error
}
