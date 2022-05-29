package internal

import (
	"log"
	"runtime/debug"
)

func Fatal(err error) {
	debug.PrintStack()
	log.Fatal(err)
}
