package fs_stringhandle

import (
	"testing"
)

func TestFirstStringIndex(t *testing.T) {
	testString := "你好fsk123/你好/log.log"

	t.Log(FirstStringIndex(testString, "/"))
}
