package menus

import (
	"context"
	"go-cms/meta"
	"go-cms/services"
)

// List get menu tree
func List(ctx context.Context) ([]meta.SimpleTreeNode, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.MenuList()
	return data, err
}

// New create menu node
func New(ctx context.Context, name, icon, url string, parentid int) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.CreateMenu(parentid, name, url, icon)
	return err == nil, err
}
