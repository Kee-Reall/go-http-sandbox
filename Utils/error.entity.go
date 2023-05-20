package errorResponse

import (
	"strings"
	"unicode"
)

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CustomErrors struct {
	ErrorsMessages []ErrorMessage `json:"errorsMessages"`
}

func formatField(incomeField string) string {
	incomeField = strings.TrimPrefix(strings.Trim(incomeField, "' "), "Key: ")
	split := strings.Split(incomeField, ".")
	runes := []rune(split[1])
	//1st letter to lowerCase
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func parseValidatorErrorString(ginMessage string) []ErrorMessage {
	split := strings.Split(ginMessage, "\n")
	var stringPairs [][]string
	for _, s := range split {
		parsed := strings.Split(s, "Error:")
		parsed[0] = formatField(parsed[0])
		stringPairs = append(stringPairs, parsed)
	}
	return newErrorMessages(stringPairs)
}

func newErrorMessage(field, message string) ErrorMessage {
	return ErrorMessage{
		field, message,
	}
}

func newErrorMessages(slice [][]string) []ErrorMessage {
	var errorMessages []ErrorMessage
	for _, value := range slice {
		errorMessages = append(errorMessages, newErrorMessage(value[0], value[1]))
	}
	return errorMessages
}

func New(er error) CustomErrors {
	var message string = er.Error()
	parsed := parseValidatorErrorString(message)
	return CustomErrors{
		parsed,
	}
}
