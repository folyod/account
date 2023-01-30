package errors

type CriticalError struct {
	message string
}

func (e *CriticalError) Error() string {
	return e.message
}
