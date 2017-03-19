package errors

func OSError(message string) ApplicationError {
	return &BaseApplicationError{
		Code:    "application.os.error",
		Message: message,
	}
}

func InputError(message string) ApplicationError {
	return &BaseApplicationError{
		Code:    "application.input.error",
		Message: message,
	}
}
