package app

import (
	"fmt"
)

type ErrTodo struct {
	msg string
}

func (e ErrTodo) Error() string {
	return fmt.Sprintf("todo error %s", e.msg)
}

type errNotFound struct {
	msg string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("todo with id %s not found ", e.msg)
}

// ErrNotFound returns the err Not Found
func ErrNotFound(msg string) errNotFound {
	return errNotFound{msg}
}
