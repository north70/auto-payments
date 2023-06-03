package tg_client

type MessageValidationError struct {
	Message string
}

func NewMessageValidationError(message string) *MessageValidationError {
	return &MessageValidationError{Message: message}
}

func (err *MessageValidationError) Error() string {
	if err.Message == "" {
		return "Ошибка обработки сообщения, повторите ещё раз"
	}
	return err.Message
}
