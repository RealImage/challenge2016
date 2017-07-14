package err /******** AUTHOR: NAGA SAI AAKARSHIT BATCHU ********/

import "fmt"

//CustomError is a Custom Error Struct
type CustomError struct {
	Type    string
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v: %v", e.Type, e.Message)
}

//OsError is a Custom Error method for OS Errors
func OsError(message string) error {
	return &CustomError{
		Type:    "server.os.error",
		Message: message,
	}
}

//ReadError is a Custom Error method for File Read Errors
func ReadError(message string) error {
	return &CustomError{
		Type:    "server.read.error",
		Message: message,
	}
}

//UpdateError is a Custom Error method for Update Errors
func UpdateError(message string) error {
	return &CustomError{
		Type:    "server.update.error",
		Message: message,
	}
}

//CopyError is a Custom Error method for Copying Errors
func CopyError(message string) error {
	return &CustomError{
		Type:    "server.copyfromparent.error",
		Message: message,
	}
}

//DistributionError is a Custom Error method for Distribution Errors
func DistributionError() error {
	return &CustomError{
		Type:    "server.distributorscope.error",
		Message: "Distribution Scope Error for the Given Locations",
	}
}

//EmptyFieldError is a Custom Error method for Empty Filed Errors
func EmptyFieldError() error {
	return &CustomError{
		Type:    "server.emptyfield.error",
		Message: "Empty Fileds Not Accepted",
	}
}
