package errors

type Error struct {
	Code    int
	message string
}

func New(message string) *Error {
	instance := &Error{
		message: message,
	}
	return instance
}

//实现Error协议并且返回Error中的message
func (e *Error) Error() string {
	return e.message
}
