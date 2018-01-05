package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize"
)

type yusan struct {
	Hello string
	world int
	from  int64
	other bool
	side  float32
}

func main() {
	xlsx := excelize.NewFile()

	index := xlsx.NewSheet("newsheet")

	xlsx.SetCellValue("newsheet", "a1", "world")
	xlsx.SetCellValue("newsheet", "A2", "hello")
	xlsx.SetCellValue("sheet1", "C1", "hahaha")

	xlsx.SetActiveSheet(index)

	err := xlsx.SaveAs("new.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("complate.")
	generateXLSX()

	readXLSX()
}

func generateXLSX() {
	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		row   *xlsx.Row
		cell  *xlsx.Cell
		y     yusan
		err   error
	)

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("sheet1")
	if err != nil {
		fmt.Println(err)
	}

	t := reflect.TypeOf(y)

	row = sheet.AddRow()
	for i := 0; i < t.NumField(); i++ {
		cell = row.AddCell()
		cell.Value = t.Field(i).Name
	}

	cell = sheet.Cell(2, 2)
	cell.SetValue(123)

	buf := new(bytes.Buffer)
	err = file.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("wrote %d", len(buf.Bytes()))

	err = file.Save("anao.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("me too")
}

var fileName = "userList.xlsx"

func readXLSX() {
	xlFile, err := xlsx.OpenFile("push.xlsx")
	if err != nil {
		return
	}

	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i != 0 {
				openId := row.Cells[1].Value
				fmt.Println(i)
				fmt.Println(openId)
				if i == 22 {
					oi := strings.Split(openId, "\n")
					fmt.Println(oi)
				}
			}
		}
	}
	return
}
