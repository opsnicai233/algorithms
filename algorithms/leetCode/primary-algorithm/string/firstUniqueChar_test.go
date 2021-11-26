package string

import (
	"log"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	s := "asdfjajklfbqa"
	log.Println(FirstUniqChar(s))
}

func TestFirstUniqChar2(t *testing.T) {
	s := "asdfjajklfbqa"
	log.Println(FirstUniqChar2(s))
	//s := "abcdefg"
	//for i:=0;i<len(s);i++ {
	//	log.Printf("%d, ", s[i])
	//}
}
