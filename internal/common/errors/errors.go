package errors

type Error struct {
	StatusCode int         `json:"-"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func (err *Error) Error() string {
	return err.Message
}

var (
	ErrInternalServer = &Error{
		StatusCode: 500,
		Status:     "error",
		Message:    "Internal server error",
		Data:       nil,
	}

	ErrBadRequest = &Error{
		StatusCode: 400,
		Status:     "error",
		Message:    "Bad request",
		Data:       nil,
	}

	ErrPermissionDenied = &Error{
		StatusCode: 403,
		Status:     "error",
		Message:    "Permission denied",
		Data:       nil,
	}

	ErrNotFound = &Error{
		StatusCode: 404,
		Status:     "error",
		Message:    "Not found",
		Data:       nil,
	}

	ErrAlreadyExists = &Error{
		StatusCode: 409,
		Status:     "error",
		Message:    "Already exists",
		Data:       nil,
	}

	ErrUnauthenticated = &Error{
		StatusCode: 401,
		Status:     "error",
		Message:    "Unauthorized",
		Data:       nil,
	}
)
