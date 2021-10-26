package firestore

import (
	"context"
	"time"
	"todoApp/pkg/app"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FirestoreRepository _ _ _ _
type FirestoreRepository struct {
	*firestore.Client
}

// NewFirestoreRepository _ _ _
func NewFirestoreRepository(fs *firestore.Client) FirestoreRepository {
	if fs == nil {
		panic("no firestore client")
	}
	return FirestoreRepository{fs}
}

func (fr *FirestoreRepository) todosCollection() *firestore.CollectionRef {
	return fr.Collection("todos")
}

func (fr FirestoreRepository) GetTodo(ctx context.Context, id string) (*app.Todo, error) {
	data, err := fr.todosCollection().Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, app.ErrNotFound(id)
		}
		if err != nil {
			return nil, errors.Wrap(err, "unable to get actual docs")
		}
	}
	tm := &TodoModel{}
	err = data.DataTo(&tm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load db documents")
	}
	t, err := fr.unmarshallTodo(tm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create todo")
	}
	return t, nil
}

func (fr FirestoreRepository) CreateTodo(ctx context.Context, todo *app.Todo) error {
	tm := fr.marshallTodo(todo)

	doc := fr.todosCollection().NewDoc()
	tm.ID = doc.ID
	_, err := doc.Create(ctx, tm)
	if err != nil {
		return errors.New("unable to save to db")
	}
	todo.ID = doc.ID
	return nil
}

func (fr FirestoreRepository) UpdateTodo(ctx context.Context, id string, updatefn func(ctx context.Context, t *app.Todo) error) (*app.Todo, error) {
	doc, err := fr.todosCollection().Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, app.ErrNotFound(id)
		}
		if err != nil {
			return nil, errors.Wrap(err, "unable to get actual docs")
		}
	}

	// get doc to TodoModel
	tm := TodoModel{}
	err = doc.DataTo(&tm)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load db document")
	}

	// map tm to app.Todo
	t, err := fr.unmarshallTodo(&tm)
	if err != nil {
		return nil, err
	}

	//  update app.Todo using logic defined by application use-cases
	if err := updatefn(ctx, t); err != nil {
		return nil, errors.Wrap(err, "unable to apply update")
	}

	// map app.Todo back to tm
	tm = fr.marshallTodo(t)
	tm.UpdatedAt = time.Now()

	// save tm to db
	if _, err = fr.todosCollection().Doc(id).Set(ctx, tm); err != nil {
		return nil, errors.Wrap(err, "unable to update db")
	}
	return t, nil
}

func (fr FirestoreRepository) DeleteTodo(ctx context.Context, id string, deletefn func(context.Context, *app.Todo)) error {
	doc, err := fr.todosCollection().Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return app.ErrNotFound(id)
		}
		if err != nil {
			return errors.Wrap(err, "unable to get actual docs")
		}
	}

	// get doc to TodoModel
	tm := TodoModel{}
	err = doc.DataTo(&tm)
	if err != nil {
		return errors.Wrap(err, "unable to load db document")
	}

	// map tm to app.Todo
	t, err := fr.unmarshallTodo(&tm)
	if err != nil {
		return err
	}
	// make struct empty
	deletefn(ctx, t)
	// Alternative approach - Instead of delete, mark deletedAt field with a data

	// delete in db
	if _, err = fr.todosCollection().Doc(id).Delete(ctx); err != nil {
		return errors.Wrap(err, "unable to delete db")
	}
	return nil
}
func (fr FirestoreRepository) unmarshallTodo(tm *TodoModel) (*app.Todo, error) {
	if tm == nil {
		return nil, errors.New("empty todo")
	}
	// if there are some custom types in domain, it can be checked or modified before saving it to db using db struct
	// Using domain logic defined in app to check for constraints
	return app.UnmarshalTodo(tm.ID,
		tm.Title,
		tm.Detail,
		tm.Done,
	)
}

func (fr FirestoreRepository) marshallTodo(t *app.Todo) TodoModel {
	tm := TodoModel{
		ID:     t.ID,
		Title:  t.String(),
		Detail: t.Detail(),
		Done:   t.Done(),
	}
	// if there some database changes, it can be effected here to be backward compatible
	return tm
}
