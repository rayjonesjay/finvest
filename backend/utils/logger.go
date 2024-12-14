package utils

import (
	"log"
	"os"
)

var Logger = log.Logger{}

func LogWriter(fd *os.File) {
	defer fd.Close()
	Logger.SetOutput(fd)
}
