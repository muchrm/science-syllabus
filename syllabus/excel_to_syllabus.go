package syllabus

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows[1:] {
		if len(row.Cells[0].String()) > 0 {
			syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", row.Cells[0].String(), row.Cells[1].String())
			_, notExist := SyllabusNotExist(db, syllabusName)
			if notExist == true {
				InsertSyllabus(
					db,
					Syllabus{
						SyllabusName: syllabusName,
					},
				)
			}
			// fmt.Printf("%s %s %s %s\n ", row.Cells[0].String(), row.Cells[1].String(), row.Cells[2].String(), row.Cells[3].String())
		} else {
			break
		}
	}
}
