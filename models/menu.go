package models

type Menu struct {
	Model
	ParentID int
	Name     string
	Icon     string
	URL      string
	Status   ENUMSTATUS
	Children []Menu
}
