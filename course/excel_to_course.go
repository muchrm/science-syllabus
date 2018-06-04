package course

import (
	"context"
	"fmt"

	"github.com/muchrm/science-syllabus/syllabus"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows[1:] {
		if len(row.Cells[0].String()) > 0 {
			courses := db.Collection("courses")
			course := Course{}
			err := courses.FindOne(
				context.Background(),
				bson.NewDocument(
					bson.EC.Int32("courseCode", stringToInt32(row.Cells[3].String())),
				),
			).Decode(course)
			if err != nil {
				fmt.Println(row.Cells[3].String())
				// panic(err)
			} else {
				syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", row.Cells[0].String(), row.Cells[1].String())
				syllabus.AddCourse(db, syllabusName, course.ID.Hex())
			}

		} else {
			break
		}
	}
}
