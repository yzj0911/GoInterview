package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

func main() {
	mysqlConn := fmt.Sprintf("%s:%v@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "123456", "192.168.1.220", 3308, "hsmstar",
	)
	var err error
	DB, err = gorm.Open(mysql.Open(mysqlConn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/GinGetSysMenu", GinGetSysMenu)
	r.Run(":3333")

}

type GetSysMenuRequest struct {
}
type GetSysMenuResponse struct {
	Nodes []*Node
}

func GinGetSysMenu(c *gin.Context) {

	var getMenuReq GetSysMenuRequest
	if err := c.Bind(&getMenuReq); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getMenuReq)
	res, err := getSysMenu(getMenuReq)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res)
}

func getSysMenu(GetSysMenuRequest) (GetSysMenuResponse, error) {
	var (
		res      GetSysMenuResponse
		err      error
		SysMenus []*SysMenu
	)
	if err = DB.Model(&SysMenu{}).Find(&SysMenus).Order("parent_id desc").Error; err != nil {
		return res, err
	}

	node := getTreeRecursive(SysMenus, 0)
	res.Nodes = node
	return res, err
}

type Node struct {
	SysMenu SysMenu `json:"SysMenu"`
	Child   []*Node `json:"child"`
}

func getTreeRecursive(list []*SysMenu, parentId int64) []*Node {
	res := make([]*Node, 0)
	for i, v := range list {
		if v.ParentId == parentId {
			res[i].Child = getTreeRecursive(list, v.ParentId)
			n := &Node{
				SysMenu: *list[i],
				Child:   res[i].Child,
			}
			res = append(res, n)
		}
	}
	return res
}

type SysMenu struct {
	Id       int64  `gorm:"column:id;type:bigint(20) unsigned;not null" json:"Id"`              // 主键
	Title    string `gorm:"column:title;type:varchar(64);not null" json:"Title"`                // 菜单标题
	Icon     string `gorm:"column:icon;type:varchar(64)" json:"Icon"`                           // 菜单图标
	Uri      string `gorm:"column:uri;type:varchar(255)" json:"Uri"`                            // 菜单地址
	ParentId int64  `gorm:"column:parent_id;type:bigint(20) unsigned;not null" json:"ParentId"` // 上级菜单
	Sort     int    `gorm:"column:sort;type:int(11);not null" json:"Sort"`                      // 菜单排序
	Status   int    `gorm:"column:status;type:tinyint(4);not null" json:"Status"`               // 菜单状态:1正常2隐藏
	Lang     string `gorm:"column:lang;type:varchar(20);not null" json:"Lang"`                  // 语言
	IsDelete int    `gorm:"column:is_delete;type:tinyint(4);not null" json:"IsDelete"`          // 是否删除
}

func (s *SysMenu) TableName() string {
	return "sys_menu"
}

//
//CREATE TABLE `sys_menu` (
//`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`title` varchar(64) NOT NULL COMMENT '菜单标题',
//`icon` varchar(64) DEFAULT NULL COMMENT '菜单图标',
//`uri` varchar(255) DEFAULT NULL COMMENT '菜单地址',
//`parent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '上级菜单',
//`sort` int(11) NOT NULL DEFAULT '0' COMMENT '菜单排序',
//`status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单状态:1正常2隐藏',
//`lang` varchar(20) NOT NULL DEFAULT 'zh-cn' COMMENT '语言',
//`is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
//PRIMARY KEY (`id`) USING BTREE,
//KEY `name` (`title`) USING BTREE,
//KEY `parent` (`parent_id`) USING BTREE
//) ENGINE = InnoDB AUTO_INCREMENT = 37 DEFAULT CHARSET =utf8mb4
