package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"log"
	"strconv"
	"time"
)

func excelDateToDate(excelDate string) string {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return  excelTime.Add(time.Second * time.Duration(days*86400)).Format("2006-01-02 15:04:05")
}
func main() {
	//f, err := excelize.OpenFile("D:\\gopath\\src\\bi\\xlsx\\计划.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//style, err := f.NewStyle(`{"number_format": 0}`)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(style)
	//f.SetCellStyle("Sheet1","C4","C4" , style)
	//rows := f.GetRows("Sheet1")
	//for _, row := range rows {
	//	fmt.Println("row[2]: ", row[2])
	//}
	//fmt.Println(excelDateToDate(string(style)))
	f, err := excelize.OpenFile("D:\\gopath\\src\\bi\\xlsx\\计划.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f.GetSheetMap())
	r,err:=f.Rows("Sheet1")
	fmt.Println(r)
	//style, err := f.NewStyle(`{"number_format": 0}`)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//f.SetCellStyle("Sheet1", "D2", "D10000", style)
	//rows := f.GetRows("Sheet1")
	//for _, row := range rows {
	//	if len(row[3]) <= 0 {
	//		continue
	//	}
	//	fmt.Println(excelDateToDate(row[3]))
	//
	//}
}
