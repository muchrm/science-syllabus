package course

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func stringToInt32(numstring string) int32 {
	num, err := strconv.Atoi(numstring)
	if err != nil {
		panic(err)
	}
	return int32(num)
}
func openCSV(name string) (*csv.Reader, func()) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ','
	return reader, func() {
		file.Close()
	}
}
func FromCSV(db *mongo.Database, name string) {
	reader, closeCSV := openCSV(name)
	defer closeCSV()
	reader.Read()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}
		if CourseNotExist(db, stringToInt32(record[11])) {
			InsertCourse(db, Course{
				CourseCode:    stringToInt32(record[11]),
				CourseName:    record[12],
				CourseNameEng: record[13],
				CoursesCredit: stringToInt32(record[14]),
				CoursesPeriod: CoursesPeriod{
					Lecture:   stringToInt32(record[15]),
					Lab:       stringToInt32(record[16]),
					SelfStudy: stringToInt32(record[17]),
				},
			})
		}
	}

}
