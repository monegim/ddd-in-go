package ch03

import "github.com/google/uuid"

type someEntity struct {
	id uuid.UUID
}

func NewSomeEntity() *someEntity {
	id := uuid.New()
	return &someEntity{id: id}
}
