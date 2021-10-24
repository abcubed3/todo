package firestore

import "time"

// TodoModel reflects the structure if the db
// it should be backward compatibility
type TodoModel struct {
	Title     string     `firestore:"title"`
	Detail    string     `firestore:"detail,omitempty"`
	Done      bool       `firestore:"done"`
	ID        string     `firestore:"id"`
	CreatedAt time.Time  `firestore:"createdAt,serverTimestamp"`
	UpdatedAt time.Time  `firestore:"updatedAt,omitempty"`
	DeletedAt *time.Time `firestore:"deletedAt,omitempty"`
}
