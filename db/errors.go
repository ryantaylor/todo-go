package db

type DuplicateError struct {
	Message string
}

func (err *DuplicateError) Error() string {
	return err.Message
}

type NotFoundError struct {
	Message string
}

func (err *NotFoundError) Error() string {
	return err.Message
}
