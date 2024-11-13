package responses

var (
	DefaultSuccessResponse = General{
		StatusCode: 200,
		Status:     "success",
		Message:    "Success",
		Data:       nil,
	}

	DefaultErrorResponse = General{
		StatusCode: 500,
		Status:     "error",
		Message:    "Internal server error",
		Data:       nil,
	}
)
