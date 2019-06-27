package roles

import (
	"context"
	"encoding/json"

	"github.com/caicloud/nirvana/log"
	"github.com/narrowizard/nirvana-cms/meta"
	"github.com/narrowizard/nirvana-cms/models"
	"github.com/narrowizard/nirvana-cms/services"
)

// List returns role list
func List(ctx context.Context, search string) ([]models.Role, error) {
	var roleService = services.NewRoleService()
	return roleService.List(search)
}

// Info returns role info
func Info(ctx context.Context, roleid int) (*models.RoleInfo, error) {
	var roleService = services.NewRoleService()
	return roleService.Info(roleid)
}

// New create role
func New(ctx context.Context, name, menus string) (*models.Role, error) {
	var menuIDs = make([]int, 0)
	var err = json.Unmarshal([]byte(menus), &menuIDs)
	if err != nil {
		log.Error(err)
		return nil, meta.UnexpectedParamError.Error("menus")
	}
	var roleService = services.NewRoleService()
	return roleService.New(name, menuIDs)
}

// Delete remove a role.
func Delete(ctx context.Context, roleid int) error {
	var roleService = services.NewRoleService()
	return roleService.Update(roleid, "", nil, int(models.STATUSDELETED))
}

// Update a role
func Update(ctx context.Context, name, menus string, roleid int) (bool, error) {
	var menuIDs = make([]int, 0)
	var err = json.Unmarshal([]byte(menus), &menuIDs)
	if err != nil {
		log.Error(err)
		return false, meta.UnexpectedParamError.Error("menus")
	}
	var roleService = services.NewRoleService()
	err = roleService.Update(roleid, name, menuIDs, 0)
	return err == nil, err
}
