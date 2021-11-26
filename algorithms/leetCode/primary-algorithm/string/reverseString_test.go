package string

import (
	"log"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"
	sb := []byte(s)
	ReverseString(sb)
	log.Println(string(sb))
}
