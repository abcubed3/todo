package app

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type TodoService struct {
	DB
}

func NewTodoService(db DB) TodoService {
	return TodoService{db}
}

type DB interface {
	GetTodo(ctx context.Context, id string) (*Todo, error)
	// GetTodos(ctx context.Context, ids ...string) (Todos, error)
	CreateTodo(ctx context.Context, todo *Todo) error
	UpdateTodo(ctx context.Context, id string, updatefn func(ctx context.Context, todo *Todo) error) (*Todo, error)
	DeleteTodo(ctx context.Context, id string, deletefn func(context.Context, *Todo)) error
}

// Create and define business and application logic here
// Use cases and business mappings are here
func (ts TodoService) GetTodo(ctx context.Context, id string) (*Todo, error) {
	todo, err := ts.DB.GetTodo(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, ErrTodoNotFound
	}
	return todo, nil
}

// func (ts TodoService) GetTodos(ctx context.Context, ids ...string) (Todos, error) {
// 	return ts.DB.GetTodos(ctx, ids)
// }

func (ts TodoService) CreateTodo(ctx context.Context, todo *Todo) error {
	if err := ts.DB.CreateTodo(ctx, todo); err != nil {
		log.Println(err)
		return ErrTodoNotSaved
	}
	return nil
}

func (ts TodoService) UpdateTodo(ctx context.Context, id, title, detail string, done bool) (*Todo, error) {
	todo, err := ts.DB.UpdateTodo(ctx, id, func(ctx context.Context, todo *Todo) error {
		return todo.Update(title, detail, done)
	})
	if err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "unable to update details")
	}
	return todo, nil
}

func (ts TodoService) DeleteTodo(ctx context.Context, id string) error {
	return ts.DB.DeleteTodo(ctx, id, func(ctx context.Context, todo *Todo) {
		todo.Delete()
	})
}
