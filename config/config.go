package config

import (
	"context"
	"strconv"
	"strings"

	"github.com/mongodb/mongo-go-driver/mongo"
)

//GerPrefixs คำนำหน้าชื่อที่ใช้สำหรับตัดออกจาก อาจารย์ผู้สอน
func GetPrefixs() []string {
	return []string{
		"ผู้ช่วยศาสตราจารย์ ดร.",
		"รองศาสตราจารย์ ดร.",
		"ศาสตราจารย์ ดร.",
		"อาจารย์ ดร.",
		"ว่าที่เรือตรี",
		"ผู้ช่วยศาสตราจารย์",
		"รองศาสตราจารย์",
		"ศาสตราจารย์",
		"อาจารย์",
		"ดร.",
		"MR.",
		"MS.",
		"นาย",
		"นางสาว",
	}
}

func ConnectMongo() (*mongo.Database, func()) {
	client, err := mongo.Connect(context.Background(), GetMongoHost(), nil)
	if err != nil {
		panic(err)
	}
	db := client.Database(GetMongoDB())
	return db, func() {
		client.Disconnect(context.Background())
	}
}

//GetMainUrl คืนค่า url หลัก
func GetMainURL() string {
	return "http://10.5.1.174/registrar/"
}

//GetURL สร้างurl สำหรับ ปี เทอมปัจจุบัน
func GetURL(courseCode int, year int, semester int) string {
	URL := "http://10.5.1.174/registrar/class_info_1.asp?coursestatus=O00&facultyid=all&maxrow=1&acadyear=$year&semester=$semester&CAMPUSID=&LEVELID=&coursecode=$courseCode&coursename=&cmd=2"
	URL = strings.Replace(URL, "$year", strconv.Itoa(year), 1)
	URL = strings.Replace(URL, "$semester", strconv.Itoa(semester), 1)
	URL = strings.Replace(URL, "$courseCode", strconv.Itoa(courseCode), 1)
	return URL
}
