package models

type Menu struct {
	Model
	ParentID int
	Name     string
	Icon     string
	URL      string
	Status   ENUMSTATUS
}

func (this Menu) CID() int {
	return this.ID
}

func (this Menu) PID() int {
	return this.ParentID
}
