package string

import (
	"gopkg.in/go-playground/assert.v1"
	"strings"
	"testing"
)

func TestStrStr(t *testing.T) {
	s := "hellollopo"
	//log.Println(strings.Index(s, ""))

	assert.Equal(t, StrStr(s, "ll"), strings.Index(s, "ll"))
	assert.Equal(t, StrStr("l", "ll"), strings.Index("l", "ll"))
	assert.Equal(t, StrStr("mississippi", "issip"), strings.Index("mississippi", "issip"))

}
