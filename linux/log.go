package linux

import (
	"log"
	"strings"
)

var logger = stubLogger{}

type stubLogger struct{}

func (s stubLogger) Info(strs ...string) {
	log.Println(strings.Join(strs, " : "))
}
