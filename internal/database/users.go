package database // import "github.com/jacekk/go-rest-api-playground/internal/database

import (
	"fmt"
)

func GetUser(id uint64) (*UserAccount, error) {
	var user UserAccount
	result := DB.First(&user, id)

	if result.Error != nil {
		return nil, fmt.Errorf("User account with id '%d' was NOT found.", id)
	}

	return &user, nil
}

func GetUserByEmail(email string) (*UserAccount, error) {
	var user UserAccount
	var where = &UserAccount{Email: email}
	result := DB.Where(where).First(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("User account with email '%s' was NOT found.", email)
	}

	return &user, nil
}

func GetUserByName(name string) (*UserAccount, error) {
	var user UserAccount
	var where = &UserAccount{Name: name}
	result := DB.Where(where).First(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("User account with name '%s' was NOT found.", name)
	}

	return &user, nil
}

func GetUsers(offset uint64, limit uint64) ([]UserAccount, error) {
	var users []UserAccount
	// NOTE: offset and limit has to be before find
	result := DB.Offset(offset).Limit(limit).Find(&users)

	return users, result.Error
}

func CreateUser(user UserAccount) (*UserAccount, error) {
	byEmail, _ := GetUserByEmail(user.Email)
	byName, _ := GetUserByName(user.Name)

	// `unique_index` works but it is hard to detect which DB error occured
	if byEmail != nil {
		msg := "User account with email '%s' already exists."
		return nil, fmt.Errorf(msg, user.Email)
	}
	if byName != nil {
		msg := "User account with name '%s' already exists."
		return nil, fmt.Errorf(msg, user.Name)
	}

	result := DB.Create(&user)

	return &user, result.Error
}

func DeleteUserById(id uint64) error {
	result := DB.Where("id = ?", id).Delete(&UserAccount{})

	if result.RowsAffected == 0 {
		return fmt.Errorf("User account with id '%d' was NOT found.", id)
	}

	return result.Error
}
