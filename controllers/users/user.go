package users

import (
	"context"
	"go-cms/models"
	"go-cms/services"
)

// List get user list
func List(ctx context.Context, page int, pagesize int, search string) (*models.PagableData, error) {
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

// Delete delete user
func Delete(ctx context.Context, userid int) error {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateUser(userid, models.STATUSDELETED)
	return err
}

// Update update user info
func Update(ctx context.Context, userid int, status int) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateUser(userid, models.ENUMSTATUS(status))
	return err == nil, err
}
