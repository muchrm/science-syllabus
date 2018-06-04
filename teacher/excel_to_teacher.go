package teacher

import (
	"fmt"
	"strings"

	"github.com/muchrm/science-syllabus/syllabus"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/muchrm/science-syllabus/config"
	"github.com/muchrm/science-syllabus/util"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows[1:] {
		if len(row.Cells[0].String()) > 0 {
			text := util.RemovePrefix(row.Cells[2].String(), config.GetPrefixs())
			name := strings.Split(text, " ")
			// fmt.Printf("%s %s %s %s\n ", row.Cells[0].String(), row.Cells[1].String(), name[0], name[1])
			teacher, notExist := TeacherNotExist(db, Teacher{OfficerName: name[0], OfficerSurname: name[1]})
			if notExist == true {
				InsertTeacher(
					db,
					Teacher{OfficerName: name[0], OfficerSurname: name[1]},
				)
				teacher, _ = TeacherNotExist(db, Teacher{OfficerName: name[0], OfficerSurname: name[1]})
			}
			// fmt.Println(teacher)
			syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", row.Cells[0].String(), row.Cells[1].String())
			syllabus.AddTeacher(db, syllabusName, teacher.ID.Hex())
		} else {
			break
		}
	}
}
