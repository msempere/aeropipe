package util

import (
	"log"
	"os"

	"github.com/andrew-d/go-termutil"
)

func PanicOnError(msg string, err error) {
	if err != nil {
		log.Printf("%v: %v", msg, err)
		panic(err)
	}
}

func IsTTY() bool {
	return termutil.Isatty(os.Stdin.Fd())
}
