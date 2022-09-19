package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

func ReadCsv(file_path string) (res [][]string) {
	file, err := os.Open(file_path)
	if err != nil {
		//logs.Errorf("open_err:", err)
		fmt.Println("open_err:", err)
		return
	}
	defer file.Close()
	// 初始化csv-reader
	reader := csv.NewReader(file)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 允许懒引号（忘记遇到哪个问题才加的这行）
	reader.LazyQuotes = true
	// 返回csv中的所有内容
	records, read_err := reader.ReadAll()
	if read_err != nil {
		//Logger.Errorf("read_err:", read_err)
		fmt.Println("read_err:", read_err)
		return
	}
	return records
}

func ReadXls(file_path string) (res [][]string) {
	if xlFile, err := xls.Open(file_path, "utf-8"); err == nil {
		fmt.Println(xlFile.Author)
		//第一个sheet
		sheet := xlFile.GetSheet(0)
		if sheet.MaxRow != 0 {
			temp := make([][]string, sheet.MaxRow)
			for i := 0; i < int(sheet.MaxRow); i++ {
				row := sheet.Row(i)
				data := make([]string, 0)
				if row.LastCol() > 0 {
					for j := 0; j < row.LastCol(); j++ {
						col := row.Col(j)
						data = append(data, col)
					}
					temp[i] = data
				}
			}
			res = append(res, temp...)
		}
	} else {
		//Logger.Errorf("open_err:", err)
		fmt.Println("open_err:", err)
	}
	return res
}

func ReadXlsx(file_path string) (res [][]string) {
	if xlFile, err := xlsx.OpenFile(file_path); err == nil {
		for index, sheet := range xlFile.Sheets {
			//第一个sheet
			if index == 0 {
				temp := make([][]string, len(sheet.Rows))
				for k, row := range sheet.Rows {
					var data []string
					for _, cell := range row.Cells {
						data = append(data, cell.Value)
					}
					temp[k] = data
				}
				res = append(res, temp...)
			}
		}
	} else {
		//Logger.Errorf("open_err:", err)
		fmt.Println("open_err:", err)
	}
	return res
}