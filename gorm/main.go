package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type Product struct {
	gorm.Model
	Code   string
	Price uint
}

func  main () {
	db, err := gorm.Open(sqlite.Open( "test.db" ), &gorm.Config{})
	if err != nil { panic ( "无法连接数据库" )  }

	// 迁移架构
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "D42" , Price: 100 })

	// 读取var product Product   db.First(&product, 1 ) // 查找具有整数主键的  产品 db.First(&product, "code = ?" , "D42" ) // 查找代码为 D42 的产品

	// 更新 - 将产品的价格更新为 200
	db.Model(&Product{}).Update( "Price" , 200 ) // 更新 - 更新多个字段  db.Model(&product).Updates(Product{Price: 200 , Code: "F42 " }) // 非零字段  db.Model(&product).Updates( map [ string ] interface {}{ "Price" : 200 , "Code" : "F42" })

	// 删除 - 删除产品
	db.Delete(&Product{}, 1 )
}