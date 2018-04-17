package models

type UserInfo struct {
	ID      int
	Account string
	Menus   []UserMenu
}
