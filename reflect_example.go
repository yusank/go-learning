package main

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

func main() {
	fmt.Println("vim-go")
	type Test struct {
		Name string `xson:"名字"`
		Age  int
	}
	list := make([]*Test, 0)
	for i := 0; i < 10; i++ {
		t := &Test{
			Name: fmt.Sprintf("name%d", i),
			Age:  i,
		}

		list = sppend(list, t)
	}
	GeneralXlsxGenerator(list)
}

/*
	GeneralXlsxGenerator 通用的生成 xlsx 文件的方法
	传入的参数为 []*SomeStruct 是一组结构体地址的切片
	xlsx 文件的每一列的列头名，取结构体 field 的 tag 中的 xson 值，如未设置，则取结构体的 field 名
	如：
	type SomeStruct struct {
		Name string `xson:"名字"
		Age  int
	}
	var doc []*SomeStruct

	以上结构体最终生成的 xlsx 文件中，生成的两列的列明分别是 "名字", "Age"
*/
func GeneralXlsxGenerator(doc interface{}) (body *bytes.Buffer, err error) {
	switch reflect.TypeOf(doc).Kind() {
	case reflect.Slice:
		v := reflect.ValueOf(doc)

		var (
			xFile *xlsx.File
			sheet *xlsx.Sheet
			row   *xlsx.Row
			cell  *xlsx.Cell
		)
		xFile = xlsx.NewFile()
		sheet, err = xFile.AddSheet("sheet1")
		if err != nil {
			return
		}

		for i := 0; i < v.Len(); i++ {
			if i == 0 {
				row = sheet.AddRow()
				tt := reflect.TypeOf(v.Index(i).Elem().Interface())
				for j := 0; j < tt.NumField(); j++ {
					val, ok := tt.Field(j).Tag.Lookup("xson")
					if !ok {
						val = tt.Field(j).Name
					}
					cell = row.AddCell()
					cell.SetValue(val)
				}
			}

			row = sheet.AddRow()
			tt := reflect.TypeOf(v.Index(i).Elem().Interface())
			vv := reflect.ValueOf(v.Index(i).Elem().Interface())
			for k := 0; k < tt.NumField(); k++ {
				var a interface{}
				a = vv.Field(k).Interface()
				cell = row.AddCell()
				cell.SetValue(a)
			}
		}

		body = new(bytes.Buffer)
		err = xFile.Write(body)

	default:
		err = errors.New("invalid type")
	}

	return
}
