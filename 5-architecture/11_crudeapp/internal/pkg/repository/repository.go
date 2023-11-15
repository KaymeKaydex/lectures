package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/ds"
)

type Repository struct {
	usersEmailIndex map[string]*ds.User
}

var usersEmailIndex map[string]*ds.User = map[string]*ds.User{
	"maksim@mail.ru": &ds.User{
		UUID:  uuid.New(),
		Email: "maksim@mail.ru",
		Age:   10,
	},
	"kirill@mail.ru": &ds.User{
		UUID:  uuid.New(),
		Email: "kirill@mail.ru",
		Age:   10,
	}}

func New(ctx context.Context) (*Repository, error) {
	return &Repository{
		usersEmailIndex: usersEmailIndex,
	}, nil
}
