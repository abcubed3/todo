package app

import (
	"github.com/pkg/errors"
)

type Todo struct {
	ID     string
	Title  string
	detail string
	done   bool
}
type Todos []Todo

func NewTodo(title string) *Todo {
	if title == "" {
		return &Todo{Title: "New Todo"}
	}
	return &Todo{Title: title}
}

func NewTodoFromString(title, detail string, done bool) *Todo {
	if title == "" {
		return &Todo{Title: "New Todo"}
	}
	return &Todo{Title: title, detail: detail, done: done}
}
func New() *Todo {
	return &Todo{}
}

func (t *Todo) Done() bool {
	return t.done
}

func (t *Todo) isDone() bool {
	return !t.done
}

func (t *Todo) String() string {
	return t.Title
}

func (t Todo) Detail() string {
	if t.detail == "" {
		return ""
	}
	return t.detail
}

func (t *Todo) UpdateDetail(detail string) error {
	if len(detail) >= 1000 {
		return errors.New("details longer than 1000 characters")
	}
	t.detail = detail
	return nil
}

func (t *Todo) Update(title, detail string, done bool) error {
	if len(detail) >= 1000 {
		return errors.New("details longer than 1000 characters")
	}
	t.Title = title
	t.done = done
	t.detail = detail
	return nil
}

// UnmarshalTodo or EncodeTodo
func UnmarshalTodo(id, title, detail string, done bool) (t *Todo, err error) {
	t = NewTodo(title)
	if err := t.UpdateDetail(detail); err != nil {
		return nil, err
	}
	t.ID = id
	t.isDone()
	return
}
