package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
}
func Debug(msg string) {
	Logger.Printf("[DEBUG] %s", msg)
}

func Info(msg string) {
	Logger.Printf("[Info] %s", msg)
}

func Warning(msg string) {
	Logger.Printf("[WARNING] %s", msg)
}
func Error(msg string) {
	Logger.Printf("[ERROR] %s", msg)
}
