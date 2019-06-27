package users

import (
	"context"

	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"
	"github.com/narrowizard/nirvana-cms/services"
)

// Info get user info
func Info(ctx context.Context, id int) (*models.UserInfo, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.UserInfo(id)
	return data, err
}

// List get user list
func List(ctx context.Context, page int, pagesize int, search string) (*models.PagableData, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.UserList(page, pagesize, search)
	return data, err
}

// New create user
func New(ctx context.Context, account, password string, roleid int) (*models.User, error) {
	var managerService = services.NewManagerService()
	u, err := managerService.CreateUser(account, password, roleid)
	return u, err
}

// Delete delete user
func Delete(ctx context.Context, userid int) error {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateUser(userid, models.STATUSDELETED, 0)
	return err
}

// Update update user info
func Update(ctx context.Context, userid, status, roleid int) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateUser(userid, models.ENUMSTATUS(status), roleid)
	return err == nil, err
}

// MenuList returns user menu list
func MenuList(ctx context.Context, nirvanacmsuserid int) ([]meta.SimpleTreeNode, error) {
	var userService = services.NewUserService()
	return userService.MenuList(nirvanacmsuserid)
}
