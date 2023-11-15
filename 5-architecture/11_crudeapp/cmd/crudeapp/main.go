package main

import (
	"context"
	"errors"
	"os"

	"github.com/sirupsen/logrus"

	app "github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/app/crudeapp"
)

func main() {
	ctx := context.Background()

	a, err := app.New(ctx)
	if err != nil {
		logrus.Error("cant create new application: %s", err)

		os.Exit(2)
	}

	err = a.Run()
	if !errors.Is(err, context.Canceled) {
		logrus.Error("error run application: %s", err)

		os.Exit(2)
	}
}
