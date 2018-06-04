package teacher

import (
	"fmt"
	"log"
	"strings"

	"github.com/muchrm/science-syllabus/syllabus"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/muchrm/science-syllabus/config"
	"github.com/muchrm/science-syllabus/util"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for index, row := range sheet.Rows[1:] {
		log.Printf("insert teacher row:%d\n", index)
		if len(row.Cells[0].String()) > 0 {
			text := util.RemovePrefix(row.Cells[2].String(), config.GetPrefixs())
			name := strings.Split(text, " ")
			teacher, notExist := TeacherNotExist(db, Teacher{OfficerName: name[0], OfficerSurname: strings.Join(name[1:], " ")})
			if notExist == true {
				log.Println(fmt.Sprintf("'%s' '%s' not a teacher", name[0], name[1]))
			} else {
				syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", row.Cells[0].String(), row.Cells[1].String())
				syllabus.AddTeacher(db, syllabusName, teacher.ID.Hex())
			}
		} else {
			break
		}
	}
}
