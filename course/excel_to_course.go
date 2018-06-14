package course

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"

	"github.com/muchrm/science-syllabus/config"
	"github.com/muchrm/science-syllabus/syllabus"
	"github.com/muchrm/science-syllabus/util"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tealeg/xlsx"
)

func FromExcel(db *mongo.Database, sheet *xlsx.Sheet) {
	for _, row := range sheet.Rows[1:] {
		// log.Printf("insert course row:%d\n", index)
		if len(row.Cells[0].String()) > 0 {
			course, notExist := CourseNotExist(db, util.StringToInt(row.Cells[3].String()))
			if notExist == true {
				courseCode, _ := row.Cells[3].Int()
				var doc *goquery.Document
				for _, year := range []int{2554, 2555, 2556, 2557, 2558, 2559, 2560} {
					for _, semester := range []int{1, 2, 3} {
						url := config.GetURL(courseCode, year, semester)
						doc = util.GetDocumentFromURL(url)
						if CourseCanUse(doc) {
							break
						}
					}
				}
				if CourseCanUse(doc) {
					courseDetails := doc.Find("tr.normaldetail").First().Find("td")
					c := Course{}
					c.AddCourseCode(courseDetails.Slice(1, 2))
					c.AddCourseNameEng(courseDetails.Slice(2, 3))
					c.AddCreditAndPeriods(courseDetails.Slice(3, 4))
					c.GetCourseName(courseDetails.Slice(1, 2))
					InsertCourse(db, c)
					course, _ = CourseNotExist(db, c.CourseCode)
				}
			}
			if course != nil {
				syllabusName := fmt.Sprintf("หลักสูตร%s สาขาวิชา%s", row.Cells[0].String(), row.Cells[1].String())
				syllabus.AddCourse(db, syllabusName, course.ID.Hex())
			} else {
				log.Printf("course %s not found on reg 2554-2560", row.Cells[3].String())
			}

		} else {
			break
		}
	}
}
func CourseCanUse(doc *goquery.Document) bool {
	if doc.Find("tr.normaldetail").Size() == 0 {
		return false
	}
	html, err := doc.Find("tr.normaldetail").First().Html()
	if err != nil {
		return false
	}
	if len(html) == 0 {
		return false
	}
	return true
}
