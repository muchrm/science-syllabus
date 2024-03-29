package syllabus

import (
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/muchrm/science-syllabus/util"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for index, row := range sheet.Rows[1:] {
		log.Printf("insert syllabus row:%d\n", index)
		if len(row.Cells[0].String()) > 0 {
			syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", util.StripSpace(row.Cells[0].String()), util.StripSpace(row.Cells[1].String()))
			_, notExist := SyllabusNotExist(db, syllabusName)
			if notExist == true {
				InsertSyllabus(
					db,
					Syllabus{
						Name: syllabusName,
					},
				)
			}
		} else {
			break
		}
	}
}
