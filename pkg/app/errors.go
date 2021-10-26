package app

import (
	"fmt"

	"github.com/pkg/errors"
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

var (
	ErrTitleLong    = errors.New("title is too long")
	ErrDetailLong   = errors.New("details longer than 1000 characters")
	ErrTodoNotFound = errors.New("todo not found")
	ErrTodoNotSaved = errors.New("unable to create todo")
)
