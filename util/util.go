package util

import "strings"

//RemovePrefix ลบคำออกจากข้อความ
func RemovePrefix(word string, prefixs []string) string {
	for _, prefix := range prefixs {
		word = strings.Replace(word, prefix, "", -1)
	}
	return word
}
