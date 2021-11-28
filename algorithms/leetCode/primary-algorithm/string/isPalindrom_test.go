package string

import (
	"gopkg.in/go-playground/assert.v1"
	"log"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	//s := "A man, a plan, a canal: Panama"
	//assert.Equal(t, IsPalindrome(s), true)

	ss := "Marge, let's \"[went].\" I await {news} telegram."
	assert.Equal(t, IsPalindrome(ss), true)

	log.Println(int('0'))
}

func TestIsPalindrome2(t *testing.T) {
	ss := "Marge, let's \"[went].\" I await {news} telegram."
	assert.Equal(t, IsPalindrome2(ss), true)

	s := "A man, a plan, a canal: Panama"
	assert.Equal(t, IsPalindrome2(s), true)
}
