// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"log"
	"testing"
)

//树
type Tree struct {
	List     map[int]*Node
	Children map[int]Node
	Parents  map[int]Node
}

//节点
type Node struct {
	Id    int     `json:"id"`
	Pid   int     `json:"pid"`
	Name  string  `json:"name"`
	Child []*Node `json:"child"`
}

//将原始数据创建树结构
func (this *Tree) BuildTree(nodes []Node) {
	this.List = make(map[int]*Node, 0)
	bs, _ := json.Marshal(nodes)
	log.Println("nodes:", string(bs))
	for index, _ := range nodes {
		id := nodes[index].Id
		nodes[index].Child = make([]*Node, 0)
		this.List[id] = &nodes[index]
	}
	log.Println("list:", this.List)
	for k, _ := range this.List {
		pid := this.List[k].Pid
		if _, ok := this.List[pid]; ok {
			this.List[pid].Child = append(this.List[pid].Child, this.List[k])
		}
	}
	//取以节点1展开的树
	for k, _ := range this.List {
		if this.List[k].Id > 1 {
			delete(this.List, k)
		}
	}
}

//GetAllNode ... 获取所有子节点
func GetAllNode(node *Node) (nodes []string) {
	if len(node.Child) == 0 {
		nodes = append(nodes, node.Name)
		return nodes
	}
	for _, t := range node.Child {
		for _, n := range GetAllNode(t) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

//func main() {
//
//	//原始数据格式 目前支持转成该种方式 [{"id":1,"name":"集团总部","pid":0},{"id":3,"name":"三体集团","pid":1},{"id":2,"name":"三体有限公司","pid":1},{"id":4,"name":"有限本部","pid":2},{"id":5,"name":"集团本部","pid":3},{"id":6,"name":"三体防御","pid":2}]
//
//	//TestNode(menus)
//}

//标准格式 树生成 需要转成标准格式字段
func TestNode(t *testing.T) {
	menus := []byte(`[{"id":1,"name":"集团总部","pid":0},{"id":3,"name":"三体集团","pid":1},{"id":2,"name":"三体有限公司","pid":1},
	{"id":4,"name":"有限本部","pid":2},{"id":5,"name":"集团本部","pid":3},{"id":6,"name":"三体防御","pid":2}]
	`)
	var nodes []Node
	err := json.Unmarshal(menus, &nodes)
	if err != nil {
		log.Fatal("JSON decode error:", err)
		return
	}
	//构建树
	var exampleTree Tree
	exampleTree.BuildTree(nodes)
	bs, _ := json.Marshal(exampleTree.List)
	log.Println("tree:", string(bs))
	//获取节点1的所有子节点
	n := GetAllNode(exampleTree.List[1])
	log.Println("n:", n)

}
