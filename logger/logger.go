// Logging module to log meassages to file

package logger

import (
	"example.com/realimage_2016/constants"

	"log"
	"os"
)

type Logger struct {
	file *os.File
}

// Creates a instance with specified file
func NewLogger() (*Logger, error) {
	file, err := os.OpenFile(constants.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &Logger{file: file}, nil
}

// Writing message to the file
func (l *Logger) Log(message ...interface{}) {
	log.SetOutput(l.file)
	for _,msg := range message {
		log.Print(msg)
	}
}

// Closing the log file
func (l *Logger) Close() error {
	return l.file.Close()
}