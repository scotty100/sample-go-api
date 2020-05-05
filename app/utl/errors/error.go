package errors

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type OneHubError interface {
	HttpStatus() int
	Err() error
}

type OneHubErrorBase struct {
	err        error
	message    string
	httpStatus int
}

func (e *OneHubErrorBase) Error() string {
	return e.message
}

func (e *OneHubErrorBase) Err() error {
	return e.err
}

func (e *OneHubErrorBase) HttpStatus() int {
	return e.httpStatus
}

type BadRequestError struct {
	FieldErrors []FieldError
	*OneHubErrorBase
}

func NewOneHubError(err error, message string, httpStatus int) OneHubError {
	return &OneHubErrorBase{
		err:        err,
		message:    message,
		httpStatus: httpStatus,
	}
}

func NotFound(err error) OneHubError {
	return &OneHubErrorBase{
		err:        err,
		message:    "NOT_FOUND",
		httpStatus: 404,
	}
}

func Conflict(err error) OneHubError {
	return &OneHubErrorBase{
		err:        err,
		message:    "ALREADY_EXISTS",
		httpStatus: 409,
	}
}

func BadRequest(err error, fieldErrors []FieldError) OneHubError {

	return &BadRequestError{
		FieldErrors: fieldErrors,
		OneHubErrorBase: &OneHubErrorBase{
			err:        err,
			message:    "BAD_REQUEST",
			httpStatus: 400,
		},
	}
}