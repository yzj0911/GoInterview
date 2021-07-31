package flow

import (
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mumushuiding/util"
)

// Node represents a specific logical unit of processing and routing
// in a workflow.
// 流程中的一个节点
type Node struct {
	Name           string          `json:"name,omitempty"`
	Type           string          `json:"type,omitempty"`
	NodeID         string          `json:"nodeId,omitempty"`
	PrevID         string          `json:"prevId,omitempty"`
	ChildNode      *Node           `json:"childNode,omitempty"`
	ConditionNodes []*Node         `json:"conditionNodes,omitempty"`
	Properties     *NodeProperties `json:"properties,omitempty"`
}

// ActionConditionType 条件类型
type ActionConditionType int

const (
	// RANGE 条件类型: 范围
	RANGE ActionConditionType = iota
	// VALUE 条件类型： 值
	VALUE
)

// ActionConditionTypes 所有条件类型
var ActionConditionTypes = [...]string{RANGE: "dingtalk_actioner_range_condition", VALUE: "dingtalk_actioner_value_condition"}

// NodeType 节点类型
type NodeType int

const (
	// START 类型start
	START NodeType = iota
	ROUTE
	CONDITION
	APPROVER
	NOTIFIER
)

// ActionRuleType 审批人类型
type ActionRuleType int

const (
	MANAGER ActionRuleType = iota
	LABEL
)

// NodeTypes 节点类型
var NodeTypes = [...]string{START: "start", ROUTE: "route", CONDITION: "condition", APPROVER: "approver", NOTIFIER: "notifier"}
var actionRuleTypes = [...]string{MANAGER: "target_management", LABEL: "target_label"}

type NodeInfoType int

const (
	STARTER NodeInfoType = iota
)

var NodeInfoTypes = [...]string{STARTER: "starter"}

type ActionerRule struct {
	Type       string `json:"type,omitempty"`
	LabelNames string `json:"labelNames,omitempty"`
	Labels     int    `json:"labels,omitempty"`
	IsEmpty    bool   `json:"isEmpty,omitempty"`
	// 表示需要通过的人数 如果是会签
	MemberCount int8 `json:"memberCount,omitempty"`
	// and 表示会签 or表示或签，默认为或签
	ActType string `json:"actType,omitempty"`
	Level   int8   `json:"level,omitempty"`
	AutoUp  bool   `json:"autoUp,omitempty"`
}
type NodeProperties struct {
	// ONE_BY_ONE 代表依次审批
	ActivateType       string             `json:"activateType,omitempty"`
	AgreeAll           bool               `json:"agreeAll,omitempty"`
	Conditions         [][]*NodeCondition `json:"conditions,omitempty"`
	ActionerRules      []*ActionerRule    `json:"actionerRules,omitempty"`
	NoneActionerAction string             `json:"noneActionerAction,omitempty"`
}
type NodeCondition struct {
	Type       string `json:"type,omitempty"`
	ParamKey   string `json:"paramKey,omitempty"`
	ParamLabel string `json:"paramLabel,omitempty"`
	IsEmpty    bool   `json:"isEmpty,omitempty"`
	// 类型为range
	LowerBound      string `json:"lowerBound,omitempty"`
	LowerBoundEqual string `json:"lowerBoundEqual,omitempty"`
	UpperBoundEqual string `json:"upperBoundEqual,omitempty"`
	UpperBound      string `json:"upperBound,omitempty"`
	BoundEqual      string `json:"boundEqual,omitempty"`
	Unit            string `json:"unit,omitempty"`
	// 类型为 value
	ParamValues []string    `json:"paramValues,omitempty"`
	OriValue    []string    `json:"oriValue,omitempty"`
	Conds       []*NodeCond `json:"conds,omitempty"`
}
type NodeCond struct {
	Type  string    `json:"type,omitempty"`
	Value string    `json:"value,omitempty"`
	Attrs *NodeUser `json:"attrs,omitempty"`
}
type NodeUser struct {
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

// NodeInfo 节点信息
type NodeInfo struct {
	NodeID      string `json:"nodeId"`
	Type        string `json:"type"`
	Aprover     string `json:"approver"`
	AproverType string `json:"aproverType"`
	MemberCount int8   `json:"memberCount"`
	Level       int8   `json:"level"`
	ActType     string `json:"actType"`
}

// GetProcessConfigFromJSONFile test
func (n *Node) GetProcessConfigFromJSONFile() {
	file, err := os.Open("D:/Workspaces/go/src/github.com/go-workflow/go-workflow/processConfig2.json")
	if err != nil {
		log.Printf("cannot open file processConfig.json:%v", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(n)
	if err != nil {
		log.Printf("decode processConfig.json failed:%v", err)
	}
}
func (n *Node) add2ExecutionList(list *list.List) {
	switch n.Type {
	case NodeTypes[APPROVER], NodeTypes[NOTIFIER]:
		var aprover string
		if n.Properties.ActionerRules[0].Type == actionRuleTypes[MANAGER] {
			aprover = "主管"
		} else {
			aprover = n.Properties.ActionerRules[0].LabelNames
		}
		list.PushBack(NodeInfo{
			NodeID:      n.NodeID,
			Type:        n.Properties.ActionerRules[0].Type,
			Aprover:     aprover,
			AproverType: n.Type,
			MemberCount: n.Properties.ActionerRules[0].MemberCount,
			ActType:     n.Properties.ActionerRules[0].ActType,
		})
		break
	default:
	}
}

// IfProcessConifgIsValid 检查流程配置是否有效
func IfProcessConifgIsValid(node *Node) error {
	// 节点名称是否有效
	if len(node.NodeID) == 0 {
		return errors.New("节点的【nodeId】不能为空！！")
	}
	// 检查类型是否有效
	if len(node.Type) == 0 {
		return errors.New("节点【" + node.NodeID + "】的类型【type】不能为空")
	}
	var flag = false
	for _, val := range NodeTypes {
		if val == node.Type {
			flag = true
			break
		}
	}
	if !flag {
		str, _ := util.ToJSONStr(NodeTypes)
		return errors.New("节点【" + node.NodeID + "】的类型为【" + node.Type + "】，为无效类型,有效类型为" + str)
	}
	// 当前节点是否设置有审批人
	if node.Type == NodeTypes[APPROVER] || node.Type == NodeTypes[NOTIFIER] {
		if node.Properties == nil || node.Properties.ActionerRules == nil {
			return errors.New("节点【" + node.NodeID + "】的Properties属性不能为空，如：`\"properties\": {\"actionerRules\": [{\"type\": \"target_label\",\"labelNames\": \"人事\",\"memberCount\": 1,\"actType\": \"and\"}],}`")
		}
	}
	// 条件节点是否存在
	if node.ConditionNodes != nil { // 存在条件节点
		if len(node.ConditionNodes) == 1 {
			return errors.New("节点【" + node.NodeID + "】条件节点下的节点数必须大于1")
		}
		// 根据条件变量选择节点索引
		err := CheckConditionNode(node.ConditionNodes)
		if err != nil {
			return err
		}
	}

	// 子节点是否存在
	if node.ChildNode != nil {
		return IfProcessConifgIsValid(node.ChildNode)
	}
	return nil
}

// CheckConditionNode 检查条件节点
func CheckConditionNode(nodes []*Node) error {
	for _, node := range nodes {
		if node.Properties == nil {
			return errors.New("节点【" + node.NodeID + "】的Properties对象为空值！！")
		}
		if len(node.Properties.Conditions) == 0 {
			return errors.New("节点【" + node.NodeID + "】的Conditions对象为空值！！")
		}
		err := IfProcessConifgIsValid(node)
		if err != nil {
			return err
		}
	}
	return nil
}

// ParseProcessConfig 解析流程定义json数据
func ParseProcessConfig(node *Node, variable *map[string]string) (*list.List, error) {
	// defer fmt.Println("----------解析结束--------")
	list := list.New()
	err := parseProcessConfig(node, variable, list)
	return list, err
}
func parseProcessConfig(node *Node, variable *map[string]string, list *list.List) (err error) {
	// fmt.Printf("nodeId=%s\n", node.NodeID)
	node.add2ExecutionList(list)
	// 存在条件节点
	if node.ConditionNodes != nil {
		// 如果条件节点只有一个或者条件只有一个，直接返回第一个
		if variable == nil || len(node.ConditionNodes) == 1 {
			err = parseProcessConfig(node.ConditionNodes[0].ChildNode, variable, list)
			if err != nil {
				return err
			}
		} else {
			// 根据条件变量选择节点索引
			condNode, err := GetConditionNode(node.ConditionNodes, variable)
			if err != nil {
				return err
			}
			if condNode == nil {
				str, _ := util.ToJSONStr(variable)
				return errors.New("节点【" + node.NodeID + "】找不到符合条件的子节点,检查变量【var】值是否匹配," + str)
				// panic(err)
			}
			err = parseProcessConfig(condNode, variable, list)
			if err != nil {
				return err
			}

		}
	}
	// 存在子节点
	if node.ChildNode != nil {
		err = parseProcessConfig(node.ChildNode, variable, list)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetConditionNode 获取条件节点
func GetConditionNode(nodes []*Node, maps *map[string]string) (result *Node, err error) {
	map2 := *maps
	for _, node := range nodes {
		var flag int
		for _, v := range node.Properties.Conditions[0] {
			paramValue := map2[v.ParamKey]
			if len(paramValue) == 0 {
				return nil, errors.New("流程启动变量【var】的key【" + v.ParamKey + "】的值不能为空")
			}
			yes, err := checkConditions(v, paramValue)
			if err != nil {
				return nil, err
			}
			if yes {
				flag++
			}
		}
		// fmt.Printf("flag=%d\n", flag)
		// 满足所有条件
		if flag == len(node.Properties.Conditions[0]) {
			result = node
		}
	}
	return result, nil
}
func getConditionNode(nodes []*Node, maps *map[string]string) (result *Node, err error) {
	map2 := *maps
	// 获取所有conditionNodes
	getNodesChan := func() <-chan *Node {
		nodesChan := make(chan *Node, len(nodes))
		go func() {
			// defer fmt.Println("关闭nodeChan通道")
			defer close(nodesChan)
			for _, v := range nodes {
				nodesChan <- v
			}
		}()
		return nodesChan
	}

	//获取所有conditions
	getConditionNode := func(nodesChan <-chan *Node, done <-chan interface{}) <-chan *Node {
		resultStream := make(chan *Node, 2)
		go func() {
			// defer fmt.Println("关闭resultStream通道")
			defer close(resultStream)
			for {
				select {
				case <-done:
					return
				case <-time.After(10 * time.Millisecond):
					fmt.Println("Time out.")
				case node, ok := <-nodesChan:
					if ok {
						// for _, v := range node.Properties.Conditions[0] {
						// 	conStream <- v
						// 	fmt.Printf("接收 condition:%s\n", v.Type)
						// }
						var flag int
						for _, v := range node.Properties.Conditions[0] {
							// fmt.Println(v.ParamKey)
							// fmt.Println(map2[v.ParamKey])
							paramValue := map2[v.ParamKey]
							if len(paramValue) == 0 {
								log.Printf("key:%s的值为空\n", v.ParamKey)
								// nodeAndErr.Err = errors.New("key:" + v.ParamKey + "的值为空")
								break
							}
							yes, err := checkConditions(v, paramValue)
							if err != nil {
								// nodeAndErr.Err = err
								break
							}
							if yes {
								flag++
							}
						}
						// fmt.Printf("flag=%d\n", flag)
						// 满足所有条件
						if flag == len(node.Properties.Conditions[0]) {
							// fmt.Printf("flag=%d\n,send node:%s\n", flag, node.NodeID)
							resultStream <- node
						} else {
							// fmt.Println("条件不完全满足")
						}
					}
				}
			}
		}()
		return resultStream
	}
	done := make(chan interface{})
	// defer fmt.Println("结束所有goroutine")
	defer close(done)
	nodeStream := getNodesChan()
	// for i := len(nodes); i > 0; i-- {
	// 	getConditionNode(resultStream, nodeStream, done)
	// }
	resultStream := getConditionNode(nodeStream, done)
	// for node := range resultStream {
	// 	return node, nil
	// }
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Time out")
			return
		case node := <-resultStream:
			// result = node
			return node, nil
		}
	}
	// setResult(resultStream, done)
	// time.Sleep(1 * time.Second)
	// log.Println("----------寻找节点结束--------")
	// return result, err
}
func checkConditions(cond *NodeCondition, value string) (bool, error) {
	// 判断类型
	switch cond.Type {
	case ActionConditionTypes[RANGE]:
		val, err := strconv.Atoi(value)
		if err != nil {
			return false, err
		}
		if len(cond.LowerBound) == 0 && len(cond.UpperBound) == 0 && len(cond.LowerBoundEqual) == 0 && len(cond.UpperBoundEqual) == 0 && len(cond.BoundEqual) == 0 {
			return false, errors.New("条件【" + cond.Type + "】的上限或者下限值不能全为空")
		}
		// 判断下限，lowerBound
		if len(cond.LowerBound) > 0 {
			low, err := strconv.Atoi(cond.LowerBound)
			if err != nil {
				return false, err
			}
			if val <= low {
				// fmt.Printf("val:%d小于lowerBound:%d\n", val, low)
				return false, nil
			}
		}
		if len(cond.LowerBoundEqual) > 0 {
			le, err := strconv.Atoi(cond.LowerBoundEqual)
			if err != nil {
				return false, err
			}
			if val < le {
				// fmt.Printf("val:%d小于lowerBound:%d\n", val, low)
				return false, nil
			}
		}
		// 判断上限,upperBound包含等于
		if len(cond.UpperBound) > 0 {
			upper, err := strconv.Atoi(cond.UpperBound)
			if err != nil {
				return false, err
			}
			if val >= upper {
				return false, nil
			}
		}
		if len(cond.UpperBoundEqual) > 0 {
			ge, err := strconv.Atoi(cond.UpperBoundEqual)
			if err != nil {
				return false, err
			}
			if val > ge {
				return false, nil
			}
		}
		if len(cond.BoundEqual) > 0 {
			equal, err := strconv.Atoi(cond.BoundEqual)
			if err != nil {
				return false, err
			}
			if val != equal {
				return false, nil
			}
		}
		return true, nil
	case ActionConditionTypes[VALUE]:
		if len(cond.ParamValues) == 0 {
			return false, errors.New("条件节点【" + cond.Type + "】的 【paramValues】数组不能为空，值如：'paramValues:['调休','年假']")
		}
		for _, val := range cond.ParamValues {
			if value == val {
				return true, nil
			}
		}
		// log.Printf("key:" + cond.ParamKey + "找不到对应的值")
		return false, nil
	default:
		str, _ := util.ToJSONStr(ActionConditionTypes)
		return false, errors.New("未知的NodeCondition类型【" + cond.Type + "】,正确类型应为以下中的一个:" + str)
	}
}
