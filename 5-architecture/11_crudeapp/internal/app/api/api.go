package api

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/atomic"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/ds"
	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/service"
)

const (
	// timeouts
	apiHTTPDefaultTimeout = 30 * time.Second
	sysHTTPDefaultTimeout = 5 * time.Minute
)

type RouterBucket struct {
	AppInfo  *ds.AppInfo
	AppReady *atomic.Bool
}

type Endpoints struct {
	ctx context.Context

	sysHTTPSrv *http.Server
	httpSrv    *http.Server

	bucket RouterBucket
}

func New(ctx context.Context, src *service.Service) (*Endpoints, error) {
	return nil, nil
}

func (endpoints *Endpoints) Run() error {
	e := echo.New()

	return e.Start(":1323")
}
