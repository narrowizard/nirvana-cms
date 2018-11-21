package users

import (
	"context"
	"encoding/json"

	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"
	"github.com/narrowizard/nirvana-cms/services"

	"github.com/caicloud/nirvana/log"
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
func New(ctx context.Context, account, password, menus string) (bool, error) {
	var menuids = make([]int, 0)
	var err = json.Unmarshal([]byte(menus), &menuids)
	if err != nil {
		log.Error(err)
		return false, meta.UnexpectedParamError.Error("menus")
	}
	var managerService = services.NewManagerService()
	err = managerService.CreateUser(account, password, menuids)
	return err == nil, err
}

// Delete delete user
func Delete(ctx context.Context, userid int) error {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateUser(userid, models.STATUSDELETED, nil)
	return err
}

// Update update user info
func Update(ctx context.Context, userid, status int, menus string) (bool, error) {
	var menusData = make([]int, 0)
	var err = json.Unmarshal([]byte(menus), &menusData)
	if err != nil {
		log.Error(err)
		return false, meta.UnexpectedParamError.Error("menus")
	}
	var managerService = services.NewManagerService()
	err = managerService.UpdateUser(userid, models.ENUMSTATUS(status), menusData)
	return err == nil, err
}

// MenuList returns user menu list
func MenuList(ctx context.Context, nirvanacmsuserid int) ([]meta.SimpleTreeNode, error) {
	var userService = services.NewUserService()
	return userService.MenuList(nirvanacmsuserid)
}
