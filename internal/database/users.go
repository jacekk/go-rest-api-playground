package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"fmt"
)

func GetUser(id int64) (*UserAccount, error) {
	var user UserAccount
	result := DB.First(&user, id)

	if result.Error != nil {
		return nil, fmt.Errorf("User account with id '%d' was NOT found.", id)
	}

	return &user, nil
}

func GetUsers(offset uint64, limit uint64) ([]UserAccount, error) {
	var users []UserAccount
	// NOTE: offset and limit has to be before find
	result := DB.Offset(offset).Limit(limit).Find(&users)

	return users, result.Error
}

func CreateUser(user UserAccount) (UserAccount, error) {
	result := DB.Create(&user)

	return user, result.Error
}

func DeleteUserById(id int64) error {
	result := DB.Where("id = ?", id).Delete(&UserAccount{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("User account with id '%d' was NOT found.", id)
	}

	return result.Error
}
