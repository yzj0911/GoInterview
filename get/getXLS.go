package main

import (
	"log"
	"os"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

var xlsxTitle = []string{"字段1", "字段2", "字段3", "字段4"}
var cell *xlsx.Cell


// 将xls.Row指针对应的数据插入到xlsx.sheet中

func insertRowFromXls(sheet *xlsx.Sheet, rowDataPtr *xls.Row, rowColCount int) {
	row := sheet.AddRow()
	for i := 0; i < rowColCount; i++ {
		cell = row.AddCell()
		cell.Value = rowDataPtr.Col(i)
	}

}

// 将一个切片指针对应的数据插入到xlsx.sheet中

func insertRow(sheet *xlsx.Sheet, rowDataPtr *[]string) {
	row := sheet.AddRow()
	rowData := *rowDataPtr
	for _, v := range rowData {
		cell := row.AddCell()
		cell.Value = v
	}

}

// 获取xlsx.File对象的指针，如果文件路径不存在则新建一个文件，并返回其指针

func getXlsxFile(filePath string) *xlsx.File {
	var file *xlsx.File
	if _, err := os.Stat(filePath); err == nil {
		file, err = xlsx.OpenFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		file = xlsx.NewFile()
		sheet, err := file.AddSheet("sheet1")
		if err != nil {
			log.Fatal(err)
		}
		insertRow(sheet, &xlsxTitle)
	}

	return file

}