package utils

import (
	"log"

	"github.com/fatih/color"
)

type Log struct{}

func (l *Log) InfoLog(message string) {
	color.Blue("INFO \t %s | %s: %s \n", log.Ldate, log.Ltime, message)
}

func (l *Log) WarningLog(message string) {
	color.Yellow("WARNING \t %s | %s: %s \n", log.Ldate, log.Ltime, message)
}

func (l *Log) ErrorLog(err error) {
	color.Red("ERROR \t %s | %s: %v \n", log.Ldate, log.Ltime, err)
}
