package crudeapp

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/app/api"
	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/repository"
	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/service"
)

type App struct {
	// root context
	ctx context.Context
}

func New(ctx context.Context) (*App, error) {
	return &App{
		ctx: ctx,
	}, nil
}

func (a *App) Run() error {
	ctx := a.ctx
	repo, err := repository.New(ctx)
	if err != nil {
		logrus.Errorf("cant create new repo with err: %s", err)

		return err
	}

	src, err := service.New(ctx, repo)
	if err != nil {
		logrus.Errorf("cant create new service with err: %s", err)

		return err
	}

	api, err := api.New(ctx, src)
	if err != nil {
		logrus.Errorf("cant create endpoints with err: %s", err)

		return err
	}

	return api.Run()
}
