package course

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type (
	// CoursesPeriod type for period in course
	CoursesPeriod struct {
		Lecture   int32 `bson:"lecture,omitempty"`
		Lab       int32 `bson:"lab,omitempty"`
		SelfStudy int32 `bson:"selfStudy,omitempty"`
	}
	// Course struct for course
	Course struct {
		ID            objectid.ObjectID `bson:"_id,omitempty"`
		CourseCode    int32             `bson:"courseCode,omitempty"`
		CourseName    string            `bson:"courseName,omitempty"`
		CourseNameEng string            `bson:"courseNameEng,omitempty"`
		CoursesCredit int32             `bson:"coursesCredit,omitempty"`
		CoursesPeriod CoursesPeriod     `bson:"coursesPeriod,omitempty"`
	}
)

func CourseNotExist(db *mongo.Database, courseCode int32) bool {
	courses := db.Collection("courses")
	err := courses.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("courseCode", courseCode),
		),
	).Decode(bson.Document{})
	if err != nil {
		return true
	}
	return false
}
func InsertCourse(db *mongo.Database, course Course) {
	courses := db.Collection("courses")
	courses.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("courseCode", course.CourseCode),
			bson.EC.String("courseName", course.CourseName),
			bson.EC.String("courseNameEng", course.CourseName),
			bson.EC.Int32("coursesCredit", course.CoursesCredit),
			bson.EC.SubDocumentFromElements("coursesPeriod",
				bson.EC.Int32("lecture", course.CoursesPeriod.Lecture),
				bson.EC.Int32("lab", course.CoursesPeriod.Lab),
				bson.EC.Int32("selfStudy", course.CoursesPeriod.SelfStudy),
			),
		),
	)
}
