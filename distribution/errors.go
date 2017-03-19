package distribution

type ApplicationError struct {
	message   string
	errorCode string
}

func OSError(message string) ApplicationError {
	return ApplicationError{
		errorCode: "application.os.error",
		message:   message,
	}
}

func InputError(message string) ApplicationError {
	return ApplicationError{
		errorCode: "application.input.error",
		message:   message,
	}
}

func DistributionScopeError(location string) ApplicationError {
	return ApplicationError{
		errorCode: "distribution.scope.error",
		message:   "Access denied to location: " + location,
	}
}
