package tools

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// Hash 单向hash
func Hash(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if nil != err {
		return ""
	}
	return string(hash)
}

func CompareHash(hash string, rawStr string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawStr))
	if nil != err {
		fmt.Println("compare hash error:", err)
		return false
	}
	return true
}
