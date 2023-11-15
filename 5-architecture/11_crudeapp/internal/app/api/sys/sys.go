package sys

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	_ = enc.Encode(c.appInfo)
}

func (c *Controller) ReadyHandler(w http.ResponseWriter, _ *http.Request) {
	ok := c.ready.Load()

	httpStatus := http.StatusOK
	if !ok {
		httpStatus = http.StatusServiceUnavailable
	}

	w.WriteHeader(httpStatus)
	enc := json.NewEncoder(w)
	_ = enc.Encode(map[string]bool{
		"ready": ok,
	})
}
func (c *Controller) LiveHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	_ = enc.Encode(map[string]bool{
		"live": true,
	})
}
