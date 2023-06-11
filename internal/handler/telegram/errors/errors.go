package errors

const ErrorHandle string = "Ошибка обработки сообщения, повторите ещё раз"

type TgValidationError struct {
	Message string
}

func NewTgValidationError(message string) *TgValidationError {
	return &TgValidationError{Message: message}
}

func (e *TgValidationError) Error() string {
	if e.Message == "" {
		return ErrorHandle
	}
	return e.Message
}
