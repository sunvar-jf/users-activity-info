package displayparser

import (
	"fileparsemod/src/helpers"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func GetExcelBook(exelpath string, yearmap map[int]map[string]int, lastloginmap map[string]string) string {
	f, err := excelize.OpenFile(exelpath)
	if err != nil {
		return err.Error()
	}
	//set year, month and logincount data
	setMonthlysheetData(f, yearmap)
	setlastloginSheetData(f, lastloginmap)
	f.Save()
	return ""
}

func setlastloginSheetData(f *excelize.File, lastloginmap map[string]string) {
	f.NewSheet(helpers.LastLoginSheetName)
	style, _ := f.NewStyle(`{"fill":{"type":"gradient","color":["#FFFFFF","#E0EBF5"],"shading":1}}`)
	f.SetCellStyle(helpers.LastLoginSheetName, "A1", "A1", style)
	f.SetCellStyle(helpers.LastLoginSheetName, "B1", "B1", style)
	f.SetCellValue(helpers.LastLoginSheetName, "A1", "User Name")
	f.SetCellValue(helpers.LastLoginSheetName, "B1", "Last Login Time")
	index := 2
	for user, timesloggedin := range lastloginmap {
		f.SetCellValue(helpers.LastLoginSheetName, "A"+strconv.Itoa(index), user)
		f.SetCellValue(helpers.LastLoginSheetName, "B"+strconv.Itoa(index), timesloggedin)
		index++
	}
}

func setMonthlysheetData(f *excelize.File, yearmap map[int]map[string]int) {
	var ind = 0
	for year, monthmap := range yearmap {
		yearind := helpers.YearIndex
		//for second year
		if ind == 1 {
			f.SetCellValue("usersheet", helpers.SecondYearColumn+strconv.Itoa(yearind), year)
		}
		//for first year
		if ind == 0 {
			f.SetCellValue("usersheet", helpers.FirstYearColumn+strconv.Itoa(yearind), year)
		}
		//set months and year1 and year2 value
		for month, logincount := range monthmap {
			yearind = monthmappingindexer(month) + helpers.YearIndex
			if ind == 0 {
				if logincount > 0 {
					f.SetCellValue("usersheet", helpers.FirstYearColumn+strconv.Itoa(yearind), logincount)
				}
			}
			if ind == 1 {
				if logincount > 0 {
					f.SetCellValue("usersheet", helpers.SecondYearColumn+strconv.Itoa(yearind), logincount)
				}
			}
		}
		ind++
	}
}

func monthmappingindexer(month string) int {
	var ind = 0
	switch month {
	case "January":
		ind = 1
	case "February":
		ind = ind + 1
	case "March":
		ind = ind + 2
	case "April":
		ind = ind + 3
	case "May":
		ind = ind + 4
	case "June":
		ind = ind + 5
	case "July":
		ind = ind + 6
	case "August":
		ind = ind + 7
	case "September":
		ind = ind + 8
	case "October":
		ind = ind + 9
	case "November":
		ind = ind + 10
	case "December":
		ind = ind + 11
	}
	return ind
}
