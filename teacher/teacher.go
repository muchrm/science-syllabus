package teacher

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Teacher struct {
	ID              objectid.ObjectID `bson:"_id,omitempty"`
	OfficerCode     int               `bson:"officerCode,omitempty"`
	OfficerName     string            `bson:"officerName,omitempty"`
	OfficerSurname  string            `bson:"officerSurname,omitempty"`
	OfficerPosition string            `bson:"officerPosition,omitempty"`
	OfficerLogin    string            `bson:"officerLogin,omitempty"`
	MajorName       string            `bson:"majorName,omitempty"`
}

func TeacherNotExist(db *mongo.Database, teacher Teacher) (*Teacher, bool) {
	syllabus := db.Collection("teachers")
	var result Teacher
	err := syllabus.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("officerName", teacher.OfficerName),
			bson.EC.String("officerSurname", teacher.OfficerSurname),
		),
	).Decode(&result)
	if err != nil {
		return nil, true
	}
	return &result, false
}
func InsertTeacher(db *mongo.Database, teacher Teacher) {
	courses := db.Collection("teachers")
	_, err := courses.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("officerName", teacher.OfficerName),
			bson.EC.String("officerSurname", teacher.OfficerSurname),
		),
	)
	if err != nil {
		panic(err)
	}
}
