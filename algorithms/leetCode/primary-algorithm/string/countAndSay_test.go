package string

import (
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	log.Println(CountAndSay(1))
	log.Println(CountAndSay(2))
	log.Println(CountAndSay(3))
	log.Println(CountAndSay(4))
	log.Println(CountAndSay(5))
	assert.Equal(t, CountAndSay(6), "312211")
}

func TestCountAndSay2(t *testing.T) {
	log.Println(CountAndSay(4))
	assert.Equal(t, CountAndSay(6), "312211")
}
