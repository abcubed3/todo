package http

import (
	"todoApp/pkg/app"
)

// Todo is the struct  of a typical Todo response
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Done   bool   `json:"done"`
}

// PostTodo is a struct for a posting a Todo
type PostTodo struct {
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Done   bool   `json:"done,omitempty"`
}

// Todos is an array of Todo
type Todos struct {
	Todos []Todo `json:"todos,omitempty"`
}

// Error defines the http error message
// type Error struct {
// 	error string `json:"error"`
// 	msg   string `json:"msg,omitempty"`
// }

// func (e Error) Error() string {
// 	return fmt.Sprintf("http error %s", e.error)
// }

// UpdateTodo defines struct to update
type UpdateTodo PostTodo

func marshalTodo(t *app.Todo) *Todo {
	return &Todo{
		ID:     t.ID,
		Title:  t.String(),
		Detail: t.Detail(),
		Done:   t.Done()}
}

func unmarshalTodo(t *PostTodo) *app.Todo {
	return app.NewTodoFromString(t.Title, t.Detail, t.Done)
}
