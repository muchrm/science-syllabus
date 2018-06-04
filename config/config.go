package config

import (
	"context"

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
