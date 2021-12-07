package main

import (
	"fileparsemod/src/fileparser"
	"fmt"
	"os"
)

func main() {
	logfilepathargs := os.Args[1]
	excelfilepath := os.Args[2]
	fmt.Println("log file path is" + logfilepathargs + "excel file path is" + excelfilepath)
	fileparser.ParseFile(logfilepathargs, excelfilepath)
}
