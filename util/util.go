package util

import (
	"crypto/sha1"
	"encoding/json"
)

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func HashSlice(i []interface{}) string {
	b, _ := json.Marshal(i)
	h := sha1.New()
	h.Write(b)
	return string(h.Sum(nil))
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}
