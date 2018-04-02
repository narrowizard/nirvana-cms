package users

import (
	"context"
	"errors"
	"go-cms/models"
	"time"
)

// List get user list
func List(ctx context.Context, page int) (*models.Model, error) {
	if page == 2 {
		return &models.Model{
			ID:        1,
			CreatedAt: time.Now(),
		}, errors.New("some error")
	}
	return &models.Model{
		ID:        1,
		CreatedAt: time.Now(),
	}, nil
}
