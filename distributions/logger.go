package distributions /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import (
	"log"
)

//ErrorLog is a function to Log Error Messages
func ErrorLog(message string, err error) {
	log.Printf("%s: %s - %v", "Error", message, err)
}

//InfoLog is a function to Log Info Messages
func InfoLog(message string) {
	log.Printf("%s: %s", "INFO", message)
}

//CustomErrorLog is a function to Log Custom Error Messages
func CustomErrorLog(message string) {
	log.Printf("%s: %s", "Error", message)
}

// CustomLog is a function to Log Custom Messages
func CustomLog(message interface{}) {
	log.Printf("%v", message)
}
