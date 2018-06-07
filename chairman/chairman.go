package chairman

import (
	"fmt"
	"log"
	"strings"

	"github.com/muchrm/science-syllabus/config"
	"github.com/muchrm/science-syllabus/syllabus"
	"github.com/muchrm/science-syllabus/teacher"
	"github.com/muchrm/science-syllabus/util"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for index, row := range sheet.Rows[3:] {
		log.Printf("insert chairman row:%d\n", index)
		if len(row.Cells[0].String()) > 0 {
			text := util.RemovePrefix(row.Cells[1].String(), config.GetPrefixs())
			name := strings.Split(text, " ")
			fmt.Println(text, name)
			teacher, notExist := teacher.TeacherNotExist(db, teacher.Teacher{OfficerName: name[0], OfficerSurname: strings.Join(name[1:], " ")})
			if notExist == true {
				log.Println(fmt.Sprintf("'%s' '%s' not a teacher", name[0], name[1]))
			} else {
				name := util.RemovePrefix(row.Cells[2].String(), []string{" สาขาวิชา"})
				name = util.TranferSyllabusToFull(name)
				// log.Println(fmt.Sprintf("can insert %s %s", teacher.ID.Hex(), name))
				syllabus.AddChairman(db, name, teacher.ID.Hex())
			}
		} else {
			break
		}
	}
}
