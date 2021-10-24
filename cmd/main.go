package main

import (
	"context"
	"log"

	"todoApp/pkg/app"
	db "todoApp/pkg/db/firestore"
	h "todoApp/pkg/http"

	"cloud.google.com/go/firestore"
)

func main() {
	ctx := context.Background()

	dbClient, err := firestore.NewClient(ctx, "abcubed3")
	if err != nil {
		log.Fatal("Failed to Firestore")
	}

	firestoredb := db.NewFirestoreRepository(dbClient)
	todoService := app.NewTodoService(firestoredb)

	srv := h.NewServer(&todoService, "8000")
	srv.Run()
}
