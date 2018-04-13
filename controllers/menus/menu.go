package menus

import (
	"context"
	"go-cms/models"
	"go-cms/services"
)

// List get menu tree
func List(ctx context.Context) ([]models.Menu, error) {
	var managerService = services.NewManagerService()
	var data, err = managerService.MenuList()
	return data, err
}

// New create menu node
func New(ctx context.Context, name, icon, url string, parentid int) (bool, error) {
	return false, nil
}
