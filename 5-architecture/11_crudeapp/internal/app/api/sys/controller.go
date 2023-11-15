package sys

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/pkg/ds"
	"go.uber.org/atomic"
)

// Controller is a struct for sys endpoints
type Controller struct {
	// ctx is ctx with cfg
	ctx context.Context
	// app ready bool
	ready *atomic.Bool
	// appInfo is information about current built app
	appInfo *ds.AppInfo
}

func New(ctx context.Context, appReady *atomic.Bool, appInfo *ds.AppInfo) (*Controller, error) {
	if appReady == nil {
		return nil, errors.New("app ready pointer cant be zero")
	}

	if appInfo == nil {
		return nil, errors.New("app ready pointer cant be zero")
	}

	return &Controller{
		ctx:     ctx,
		ready:   appReady,
		appInfo: appInfo,
	}, nil
}
