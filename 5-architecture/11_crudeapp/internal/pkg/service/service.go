package service

import (
	"context"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/repository"
)

type Service struct {
	ctx context.Context // root ctx

	repo *repository.Repository
}

func New(ctx context.Context, repo *repository.Repository) (*Service, error) {
	return &Service{ctx: ctx, repo: repo}, nil
}

func (s *Service) GetUserArticlesByUserID() {

}
