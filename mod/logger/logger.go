package logger

import (
	"fmt"
	"os"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var blue = "\033[36m"
var yellow = "\033[33m"

func InfoMessage(message string, data ...interface{}) {
	fmt.Fprintf(os.Stdout, blue+message+reset+"\n", data...)
}

func ErrorMessage(message string, data ...interface{}) {
	fmt.Fprintf(os.Stdout, red+message+reset+"\n", data...)
}

func PendingMessage(message string, data ...interface{}) {
	fmt.Fprintf(os.Stdout, yellow+message+reset+"\n", data...)
}

func CompletedMessage(message string, data ...interface{}) {
	fmt.Fprintf(os.Stdout, green+message+reset+"\n", data...)
}
