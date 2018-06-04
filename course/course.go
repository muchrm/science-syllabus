package course

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/muchrm/science-syllabus/config"

	"github.com/muchrm/science-syllabus/util"

	"github.com/PuerkitoBio/goquery"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type (
	// CoursesPeriod type for period in course
	Periods struct {
		Lecture   int `bson:"lecture,omitempty"`
		Lab       int `bson:"lab,omitempty"`
		SelfStudy int `bson:"selfStudy,omitempty"`
	}
	// Course struct for course
	Course struct {
		ID            objectid.ObjectID `bson:"_id,omitempty"`
		CourseCode    int               `bson:"courseCode,omitempty"`
		CourseName    string            `bson:"courseName,omitempty"`
		CourseNameEng string            `bson:"courseNameEng,omitempty"`
		Credit        int               `bson:"credit,omitempty"`
		Periods       Periods           `bson:"periods,omitempty"`
	}
)

func (c *Course) AddCourseNameEng(courseNameNode *goquery.Selection) {
	courseNameNode.Find("font").Remove()
	c.CourseNameEng = util.StripFirstSpace(courseNameNode.Text(), "	")
}
func (c *Course) AddCourseCode(courseCodeNode *goquery.Selection) {
	var err error
	c.CourseCode, err = strconv.Atoi(courseCodeNode.Find("a").Text())
	if err != nil {
		panic(err)
	}

}
func (c *Course) AddCreditAndPeriods(creditPeriodsNode *goquery.Selection) {
	creditPeriods := strings.Split(creditPeriodsNode.Text(), " ")
	var err error
	c.Credit, err = strconv.Atoi(creditPeriods[0])
	if err != nil {
		panic(err)
	}

	period := strings.Replace(creditPeriods[1], "(", "", 1)
	period = strings.Replace(period, ")", "", 1)
	periods := []int{}
	for _, i := range strings.Split(period, "-") {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		periods = append(periods, j)
	}
	c.Periods = Periods{periods[0], periods[1], periods[2]}
}
func (c *Course) GetCourseName(linkNode *goquery.Selection) {
	if link, exist := linkNode.Find("a").Attr("href"); exist == true {
		URL := fmt.Sprintf("%v%v", config.GetMainURL(), link)
		doc := util.GetDocumentFromURL(URL)
		c.CourseName = doc.Find("table[class=\"normaldetail\"]").
			Slice(0, 1).
			Find("tr").
			Slice(1, 2).
			Find("td").
			Slice(1, 2).
			Text()
	}
}
func CourseNotExist(db *mongo.Database, courseCode int) (*Course, bool) {
	courses := db.Collection("courses")
	var result Course
	err := courses.FindOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("courseCode", int32(courseCode)),
		),
	).Decode(&result)
	if err != nil {
		return nil, true
	}
	return &result, false
}
func InsertCourse(db *mongo.Database, course Course) {
	courses := db.Collection("courses")
	courses.InsertOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int32("courseCode", int32(course.CourseCode)),
			bson.EC.String("courseName", course.CourseName),
			bson.EC.String("courseNameEng", course.CourseName),
			bson.EC.Int32("credit", int32(course.Credit)),
			bson.EC.SubDocumentFromElements("periods",
				bson.EC.Int32("lecture", int32(course.Periods.Lecture)),
				bson.EC.Int32("lab", int32(course.Periods.Lab)),
				bson.EC.Int32("selfStudy", int32(course.Periods.SelfStudy)),
			),
		),
	)
}
