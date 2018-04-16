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

func (this TreeData) CID() int {
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
	var treeableData = make([]TreeableData, 0)
	for _, v := range testData {
		treeableData = append(treeableData, v)
	}

	var tree = ArrayToTree(treeableData)
	t.Log(tree)
	if len(tree) != 1 {
		t.Errorf("got error root node count.")
		return
	}
	var root = tree[0]
	if root.Find(6) == nil {
		t.Errorf("cannot find child node.")
	}
	if root.Find(7) != nil {
		t.Errorf("find node not existed.")
	}
}
