package roles

import (
	"context"

	"github.com/narrowizard/nirvana-cms/models"
)

// List returns role list
func List(ctx context.Context, search string) ([]models.RoleInfo, error) {
	return nil, nil
}

// New create role
func New(ctx context.Context, name, menus string) (*models.RoleInfo, error) {
	return nil, nil
}

// Delete remove a role. TODO: users associated with the role will be rejected to login.
func Delete(ctx context.Context, roleid int) error {
	return nil
}

// Update a role
func Update(ctx context.Context, name, menus string, roleid int) (*models.RoleInfo, error) {
	return nil, nil
}
