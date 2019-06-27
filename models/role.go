package models

// Role 角色
type Role struct {
	Model
	Name   string
	Status ENUMSTATUS
}

// RoleMenu 角色权限绑定
type RoleMenu struct {
	Model
	RoleID int
	MenuID int
	Status ENUMSTATUS
}
