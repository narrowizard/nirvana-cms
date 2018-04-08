package users

import (
	"context"
	"go-cms/models"
	"go-cms/services"
)

// List get user list
func List(ctx context.Context, page int, pagesize int, search string) ([]models.User, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.UserList(page, pagesize, search)
	return data, err
}

// New create user
func New(ctx context.Context, account, password string) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.CreateUser(account, password)
	return err == nil, err
}
