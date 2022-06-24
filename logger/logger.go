package logger

import "log"

type LoggerCustome struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

var Logger *LoggerCustome

func SetLogger(l *LoggerCustome) {
	Logger = l
}
