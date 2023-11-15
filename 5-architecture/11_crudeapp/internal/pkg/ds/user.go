package ds

import "github.com/google/uuid"

type User struct {
	UUID  uuid.UUID
	Email string
	Age   int
}
