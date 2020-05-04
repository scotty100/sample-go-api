package errors

type ProcessingError struct {
	// Message is the error message that may be displayed to end users
	Message string

	IsRecoverable bool

	Error error
}

func (pe *ProcessingError) FullErrorMessage() string {

	errorMessage := pe.Message

	if pe.Error != nil {
		errorMessage = errorMessage + " : " + pe.Error.Error()
	}

	return errorMessage
}
