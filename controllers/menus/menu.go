package menus

import (
	"context"
	"nirvana-cms/meta"
	"nirvana-cms/services"
)

// List get menu tree
func List(ctx context.Context) ([]meta.SimpleTreeNode, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.MenuList()
	return data, err
}

// New create menu node
func New(ctx context.Context, name, icon, url, remarks string, parentid, ismenu int) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.CreateMenu(parentid, ismenu, name, url, icon, remarks)
	return err == nil, err
}

// Update update a menu node
func Update(ctx context.Context, name, icon, url, remarks string, id, ismenu int) (bool, error) {
	var managerService = services.NewManagerService()
	var err = managerService.UpdateMenu(id, ismenu, name, url, icon, remarks)
	return err == nil, err
}
