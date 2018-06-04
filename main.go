package main

import (
	"fmt"

	"github.com/muchrm/science-syllabus/course"
	"github.com/muchrm/science-syllabus/syllabus"
	"github.com/muchrm/science-syllabus/teacher"

	"github.com/muchrm/science-syllabus/config"
	"github.com/tealeg/xlsx"
)

func main() {
	db, closeConnection := config.ConnectMongo()
	defer closeConnection()
	// course.FromCSV(db, "./course.csv")
	xlFile, err := xlsx.OpenFile("./sheet.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	syllabus.FromExcel(db, xlFile.Sheets[0])
	teacher.FromExcel(db, xlFile.Sheets[0])
	course.FromExcel(db, xlFile.Sheets[4])
	course.FromExcel(db, xlFile.Sheets[5])
	course.FromExcel(db, xlFile.Sheets[6])
}
