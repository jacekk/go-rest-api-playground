package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"encoding/json"
)

func (self *Post) UnmarshalJSON(bytes []byte) error {
	type PostAlias Post
	var category PostCategory

	post := &struct {
		Category string
		*PostAlias
	}{
		PostAlias: (*PostAlias)(self),
	}
	err := json.Unmarshal(bytes, &post)

	if err != nil {
		return err
	}

	DB.FirstOrCreate(&category, PostCategory{Name: post.Category})
	self.CategoryID = category.ID
	self.Category = category

	return nil
}

func (self *UserAccount) UnmarshalJSON(bytes []byte) error {
	type UserAccountAlias UserAccount

	user := &struct {
		Password string
		*UserAccountAlias
	}{
		UserAccountAlias: (*UserAccountAlias)(self),
	}
	err := json.Unmarshal(bytes, &user)

	if err != nil {
		return err
	}

	self.EncryptPasswordIfSet(user.Password)

	return nil
}
