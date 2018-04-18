package models

type Menu struct {
	Model
	ParentID int
	IsMenu   int // 1-模块 2-接口
	Name     string
	Remarks  string
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
