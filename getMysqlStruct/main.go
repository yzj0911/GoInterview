package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("D:\\gopath\\src\\execlt1\\model.go").
		Dsn("root:123456@tcp(192.168.3.111:3306)/bi?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table("core_theme").
		Run()
	fmt.Println(err)
}
