package errorResponse

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CustomErrors struct {
	ErrorsMessages []ErrorMessage `json:"errorsMessages"`
}

func NewException(er error) CustomErrors {
	var message string = er.Error()
	return CustomErrors{
		ErrorsMessages: []ErrorMessage{
			ErrorMessage{"name", message},
			ErrorMessage{"websiteUrl", message},
		},
	}
}
