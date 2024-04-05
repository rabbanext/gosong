package middlewares

import (
	"errors"

	"github.com/rabbanext/gosong/models"
)

// Simulate a database call
func FindByCredentials(email, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	if email == "test@mail.com" && password == "test12345" {
		return &models.User{
			ID:             1,
			Email:          "test@mail.com",
			Password:       "test12345",
			FavoritePhrase: "Hello, World!",
		}, nil
	}
	return nil, errors.New("user not found")
}
