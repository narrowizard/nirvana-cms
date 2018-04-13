package meta

import (
	"testing"
)

type TreeData struct {
	id  int
	pid int
}

func NewTreeData(id, pid int) TreeData {
	return TreeData{
		id:  id,
		pid: pid,
	}
}

func (this TreeData) ID() int {
	return this.id
}

func (this TreeData) PID() int {
	return this.pid
}

var testData = []TreeData{
	NewTreeData(6, 2),
	NewTreeData(1, 0),
	NewTreeData(2, 1),
	NewTreeData(3, 1),
	NewTreeData(4, 1),
	NewTreeData(5, 1),
}

func TestArrayToTree(t *testing.T) {
	var data = make([]TreeableData, 0)
	for _, v := range testData {
		data = append(data, TreeableData(v))
	}
	var tree = ArrayToTree(data)
	if len(tree) != 1 {
		t.Errorf("wrong tree data: %+v", tree)
	}
}
