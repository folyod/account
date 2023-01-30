package assertions

type assertionError struct {
	code    string
	message string
	context *context
}

func newError(code string) assertionError {
	message, isExist := codemap[code]
	if !isExist {

	}

	return assertionError{
		code:    code,
		message: message,
		context: nil,
	}
}

func (e *assertionError) RewriteDefaultMessage(message string) {
	e.message = message
}

func (e *assertionError) WithContext(context *context) {
	e.context = context
}

func (e *assertionError) Code() string {
	return e.code
}

func (e *assertionError) ServiceMessage() string {
	return e.message
}

func (e *assertionError) HasContext() bool {
	return e.context != nil
}
