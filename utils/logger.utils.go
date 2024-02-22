package utils

import (
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct{}

func (l *Logger) InfoLog(message, pack string) {
	color.Blue("INFO %s | package=%s: %s \n", time.Now().String(), pack, message)
}

func (l *Logger) WarningLog(message, pack string) {
	color.Yellow("WARNING %s | package=%s: %s \n", time.Now().String(), message)
}

func (l *Logger) ErrorLog(err error, pack string) {
	color.Red("ERROR %s | package=%s: %v \n", time.Now().String(), pack, err)
}

func (l *Logger) FatalLog(err error, pack string) {
	color.Red("FATAL %s | package=%s: %v \n", time.Now().String(), pack, err)
	os.Exit(1)
}

var Log = Logger{}
