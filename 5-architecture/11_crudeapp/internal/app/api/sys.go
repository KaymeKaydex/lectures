package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/go-park-mail-ru/lectures/5-architecture/11_crudeapp/internal/app/api/sys"
)

func (endpoints *Endpoints) InitSysRoutes(ctx context.Context) error {
	var err error

	// using default http server
	mux := http.NewServeMux()

	sysController, err := sys.New(ctx, endpoints.bucket.AppReady, endpoints.bucket.AppInfo)
	if err != nil {
		return err
	}
	{
		// version
		mux.HandleFunc("/version", sysController.VersionHandler)
		// k8s probes
		mux.HandleFunc("/ready", sysController.ReadyHandler)
		mux.HandleFunc("/live", sysController.LiveHandler)
		// prometheus metrics
		mux.Handle("/metrics", promhttp.Handler())
		// pprof
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	}

	port := "8400"

	s := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: sysHTTPDefaultTimeout,
		ReadTimeout:  sysHTTPDefaultTimeout,
		IdleTimeout:  sysHTTPDefaultTimeout,
		Handler:      mux,
	}
	endpoints.sysHTTPSrv = s

	logrus.WithField("port", port).Info("sys-http server started serving")

	err = s.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		logrus.WithError(err).Error("sys-http server init/runtime error")
	}

	return err
}
