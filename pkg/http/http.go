package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
	"todoApp/pkg/app"

	"github.com/gorilla/mux"
)

type Server struct {
	*app.TodoService
	r *mux.Router
	*http.Server
}

func NewServer(t *app.TodoService, port string) Server {
	r := mux.NewRouter()
	addr := ":" + port
	// Good practice: enforce timeouts for servers you create!
	srv := &http.Server{Handler: r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second}
	return Server{t, r, srv}
}

func (s Server) get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		fmt.Fprint(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	todo, err := s.GetTodo(ctx, id)
	if err != nil {
		// log.Println(err)
		if errors.Is(err, app.ErrTodoNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	data := marshalTodo(todo)
	writeTodoJSON(w, data)
	w.WriteHeader(http.StatusOK)
}

func (s Server) create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	todo := &PostTodo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Unmarshall todo to app.Todo
	t := unmarshalTodo(todo)
	err := s.CreateTodo(ctx, t)
	if err != nil {
		log.Println(err)
		return
	}
	data := marshalTodo(t)
	w.WriteHeader(http.StatusCreated)
	writeTodoJSON(w, data)

}

func (s Server) updateDetail(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		fmt.Fprint(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	todo := &PostTodo{}
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	payload, err := s.UpdateTodo(ctx, id, todo.Title, todo.Detail, todo.Done)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := marshalTodo(payload)
	writeTodoJSON(w, data)
}

func (s Server) delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		fmt.Fprint(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err := s.DeleteTodo(ctx, id)
	if err != nil {
		// log.Println(err)
		if errors.Is(err, app.ErrTodoNotFound) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}
