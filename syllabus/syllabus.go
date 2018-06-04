package syllabus

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Syllabus struct {
	ID           objectid.ObjectID `bson:"_id,omitempty"`
	SyllabusName string            `bson:"syllabusName,omitempty"`
}

func SyllabusNotExist(db *mongo.Database, name string) (Syllabus, bool) {
	syllabus := db.Collection("syllabus")
	result := Syllabus{}
	err := syllabus.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("syllabusName", name),
		),
	).Decode(&result)
	if err != nil {
		return result, true
	}
	return result, false
}
func InsertSyllabus(db *mongo.Database, syllabus Syllabus) {
	courses := db.Collection("syllabus")
	_, err := courses.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("syllabusName", syllabus.SyllabusName),
		),
	)
	if err != nil {
		panic(err)
	}
}

func AddTeacher(db *mongo.Database, syllabusName string, teacherID string) {
	syllabus := db.Collection("syllabus")
	syllabus.FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("syllabusName", syllabusName),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push", bson.EC.String("teachers", teacherID)),
		),
	)
}
func AddCourse(db *mongo.Database, syllabusName string, courseID string) {
	syllabus := db.Collection("syllabus")
	syllabus.FindOneAndUpdate(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("syllabusName", syllabusName),
		),
		bson.NewDocument(
			bson.EC.SubDocumentFromElements("$push", bson.EC.String("courses", courseID)),
		),
	)
}
