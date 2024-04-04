package tools

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := Hash("1234")
	fmt.Println(hash)
	//
}

func TestCompareHash(t *testing.T) {
	hash := "$2a$10$nIQg3zqjAoMJMZUUA6RVq.7ruUbLKMPmgOt56lOPEJIAp2tTDL1cm"
	ok := CompareHash(hash, "1234")
	fmt.Println("pass = ", ok)
}
