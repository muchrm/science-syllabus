package main

import (
	"log"

	"github.com/muchrm/science-syllabus/chairman"
	"github.com/muchrm/science-syllabus/config"
	"github.com/tealeg/xlsx"
)

func main() {
	db, closeConnection := config.ConnectMongo()
	defer closeConnection()
	// xlFile, err := xlsx.OpenFile("./sheet.xlsx")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	xlFile2, err := xlsx.OpenFile("./sheet_2.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println("*-----------------------------------Insert syllabus-----------------------------------*")
	// syllabus.FromExcel(db, xlFile.Sheets[0])
	// log.Println("*-----------------------------------Insert teacher in syllabus-----------------------------------*")
	// teacher.FromExcel(db, xlFile.Sheets[0])

	// log.Println("*-----------------------------------Insert course 1 in syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[4])
	// log.Println("*-----------------------------------Insert course 2 syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[5])
	// log.Println("*-----------------------------------Insert course 3 syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[6])
	// log.Println("*-----------------------------------Insert course 4 syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[9])
	// log.Println("*-----------------------------------Insert course 5 syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[10])
	// log.Println("*-----------------------------------Insert course 6 syllabus-----------------------------------*")
	// course.FromExcel(db, xlFile2.Sheets[11])
	log.Println("*-----------------------------------Insert chairman 1 -----------------------------------*")
	chairman.FromExcel(db, xlFile2.Sheets[7])
	log.Println("*-----------------------------------Insert chairman 2 -----------------------------------*")
	chairman.FromExcel(db, xlFile2.Sheets[8])
}
