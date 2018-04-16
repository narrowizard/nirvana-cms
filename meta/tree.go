package meta

// TreeableData array data that could translate to tree
type TreeableData interface {
	CID() int
	PID() int
}

// SimpleTreeNode 树节点的简单实现
type SimpleTreeNode struct {
	Data     TreeableData
	Children []*SimpleTreeNode
}

func (this *SimpleTreeNode) Append(cnode *SimpleTreeNode) *SimpleTreeNode {
	this.Children = append(this.Children, cnode)
	return this
}

// Find 从当前节点与子节点查找id为key的节点
func (this *SimpleTreeNode) Find(key int) *SimpleTreeNode {
	if this.Data.CID() == key {
		return this
	}
	for _, v := range this.Children {
		var temp = v.Find(key)
		if temp != nil {
			return temp
		}
	}
	return nil
}

func NewSimpleTreeNode(data TreeableData) *SimpleTreeNode {
	var temp = &SimpleTreeNode{
		Data: data,
	}
	temp.Children = make([]*SimpleTreeNode, 0)
	return temp
}

// ArrayToTree translate array data to nested tree data
func ArrayToTree(array []TreeableData) []SimpleTreeNode {
	var data = make([]SimpleTreeNode, 0)
	// 先找出所有顶级节点
	var cids = make(map[int]bool)
	var pids = make(map[int]bool)
	for _, v := range array {
		cids[v.CID()] = true
		pids[v.PID()] = true
	}

	// 计算 pids - cids, 得到顶级节点
	var roots = make(map[int]bool)
	for p := range pids {
		if !cids[p] {
			roots[p] = true
		}
	}

	var handled = false
	for i := 0; i < len(array); i++ {
		var v = array[i]
		// if root
		if roots[v.PID()] {
			data = append(data, *NewSimpleTreeNode(v))
			continue
		}
		// find parent node
		handled = false
		for i := 0; i < len(data); i++ {
			var pnode = data[i].Find(v.PID())
			if pnode != nil {
				pnode.Append(NewSimpleTreeNode(v))
				handled = true
				break
			}
		}
		if !handled {
			// append to array end for next handle
			array = append(array, v)
		}
	}
	return data
}
