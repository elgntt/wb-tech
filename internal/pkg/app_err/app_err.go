package app_err

type BusinessError struct {
	message string
}

func (b BusinessError) Error() string {
	return b.message
}

func NewBusinessError(message string) error {
	return BusinessError{
		message: message,
	}
}
