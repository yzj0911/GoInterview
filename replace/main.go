package main

import (
	"encoding/json"
	"fmt"
)

var city = []string{
	"杭州",
	"宁波",
	"温州",
	"嘉兴",
	"湖州",
	"绍兴",
	"金华",
	"衢州",
	"舟山",
	"台州",
	"丽水",
}

func main() {
	//a := "浙江省温州市龙湾区挖的放垃圾收代理费处女座
	var str =[1]string{"浙江省杭州市龙湾区新城挖的放垃圾收代理费处女座"}
	//"	^.+县.+小区.+号楼.+单元.+(室|户).*$" java
	///.+?(省|市|自治区|自治州|县|区)/g javascript
	//([^省]+省|.+自治区|[^市]+市)([^自治州]+自治州|[^市]+市|[^盟]+盟|[^地区]+地区|.+区划)([^市]+市|[^县]+县|[^旗]+旗|.+区)
	//regex := `(杭州|宁波|温州|嘉兴|湖州|绍兴|金华|衢州|舟山|台州|丽水)`
	//s1 := regexp.MustCompile(regex)
	//fmt.Println(s1.FindString(str))
	//str := "脚本之家11"
	//matched, err := regexp.MatchString("[\u4e00-\u9fa5]", str)
	//fmt.Println(matched, err)

	a,_:=json.Marshal(str)
fmt.Println(string(a))
}
func A(as []int) {
	as[0] = 4
	fmt.Println(as)
}

var districtCountyCitys []districtCountyCity

type districtCountyCity struct {
	province string
	city     string
	county   string
}

func analysisPlace(place string) districtCountyCity {
	var d districtCountyCity
	//regex:=[]
	return d
}
