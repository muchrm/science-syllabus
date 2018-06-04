package util

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
)

func StringToInt(numstring string) int {
	num, err := strconv.Atoi(numstring)
	if err != nil {
		panic(err)
	}
	return int(num)
}

//RemovePrefix ลบคำออกจากข้อความ
func RemovePrefix(word string, prefixs []string) string {
	for _, prefix := range prefixs {
		word = strings.Replace(word, prefix, "", -1)
	}
	return word
}

//StripFirstSpace ลบช่องว่างด้านหน้าข้อความ
func StripFirstSpace(word string, match string) string {
	return strings.Replace(word, match, "", 1)
}

//StripSpace ลบช่องว่างออกจากข้อความ ในบางครั้งจะไม่สามารถลบออกได้ด้วยวิธีปกติ จึงต้องลบแบบเทียบตัวอักษร
func StripSpace(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, word)
}
func GetDocumentFromURL(URL string) *goquery.Document {
	content, err := GetBodyString(URL)
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		panic(err)
	}
	return doc
}

//GetBodyString แปลง http body เป็น format utf-8 และ คัดลอกไปยังreader
func GetBodyString(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("get content  error:", err)
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	body := strings.NewReader(buf.String())

	utfBody, err := iconv.NewReader(body, "windows-874", "utf-8")
	if err != nil {
		return nil, err
	}
	return utfBody, nil
}
