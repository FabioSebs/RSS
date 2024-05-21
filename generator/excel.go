package generator

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type ExcelGen struct {
	Controller *excelize.File
}

func NewExcelGen() ExcelGen {
	return ExcelGen{
		Controller: excelize.NewFile(),
	}
}

func (x *ExcelGen) NewSheet(name string) (sheet int, err error) {
	sheet, err = x.Controller.NewSheet(name)
	if err != nil {
		return
	}
	return
}

func (x *ExcelGen) SetColumns(sheetNo int, cols []string) error {
	for idx := 1; idx < len(cols)+1; idx++ {
		switch idx {
		case 1:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), "A1", cols[idx-1]); err != nil {
				return err
			}
		case 2:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), "B1", cols[idx-1]); err != nil {
				return err
			}
		case 3:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), "C1", cols[idx-1]); err != nil {
				return err
			}
		case 4:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), "D1", cols[idx-1]); err != nil {
				return err
			}
		default:
			fmt.Println("done")
		}
	}
	return nil
}

func (x *ExcelGen) SetValues(rowNo int, sheetNo int, values []string) error {
	for idx := 1; idx < len(values)+1; idx++ {
		switch idx {
		case 1:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), fmt.Sprintf("A%d", rowNo), values[idx-1]); err != nil {
				return err
			}
		case 2:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), fmt.Sprintf("B%d", rowNo), values[idx-1]); err != nil {
				return err
			}
		case 3:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), fmt.Sprintf("C%d", rowNo), values[idx-1]); err != nil {
				return err
			}
		case 4:
			if err := x.Controller.SetCellValue(fmt.Sprintf("Sheet%d", sheetNo), fmt.Sprintf("D%d", rowNo), values[idx-1]); err != nil {
				return err
			}
		default:
			fmt.Println("done")
		}

	}

	return nil
}

func (x *ExcelGen) SaveFile(filename string) error {
	if err := x.Controller.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
	return nil
}
