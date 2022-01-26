package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	IsOnMyself int `form:"is_on_myself" binding:"required"`
	传入0时，会报错。
	ERROR Key: 'GetInferenceTaskRule.IsOnMyself' Error:Field validation for 'IsOnMyself' failed on the 'required' tag
	修改 IsOnMyself *int `form:"is_on_myself" binding:"required"`
*/
func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.POST("/Pos", func(c *gin.Context) {
		var getInferenceTaskRule GetInferenceTaskRule
		if err := c.Bind(&getInferenceTaskRule); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(getInferenceTaskRule)
		c.JSON(http.StatusOK, getInferenceTaskRule)
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8111")
}

type GetInferenceTaskRule struct {
	Page       int `form:"page" binding:"required"`
	Size       int `form:"size" binding:"required"`
	IsOnMyself int `form:"is_on_myself" binding:"required"`
}
