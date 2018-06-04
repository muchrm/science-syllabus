package main

import (
	"log"

	"github.com/muchrm/science-syllabus/course"
	"github.com/muchrm/science-syllabus/syllabus"
	"github.com/muchrm/science-syllabus/teacher"

	"github.com/muchrm/science-syllabus/config"
	"github.com/tealeg/xlsx"
)

func main() {
	db, closeConnection := config.ConnectMongo()
	defer closeConnection()
	xlFile, err := xlsx.OpenFile("./sheet.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("*-----------------------------------Insert syllabus-----------------------------------*")
	syllabus.FromExcel(db, xlFile.Sheets[0])
	log.Println("*-----------------------------------Insert teacher in syllabus-----------------------------------*")
	teacher.FromExcel(db, xlFile.Sheets[0])
	log.Println("*-----------------------------------Insert course 1 in syllabus-----------------------------------*")
	course.FromExcel(db, xlFile.Sheets[4])
	log.Println("*-----------------------------------Insert course 2 syllabus-----------------------------------*")
	course.FromExcel(db, xlFile.Sheets[5])
	log.Println("*-----------------------------------Insert course 3 syllabus-----------------------------------*")
	course.FromExcel(db, xlFile.Sheets[6])
}
