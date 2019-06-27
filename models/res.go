package models

import "time"

type UserInfo struct {
	ID         int
	Account    string
	RoleID     int
	RoleName   string
	RoleStatus ENUMSTATUS
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     ENUMSTATUS
}

type RoleInfo struct {
	ID    int
	Name  string
	Menus []RoleMenu
}
