package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *server) HealthHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get health api called")
	writeResponse(w, http.StatusOK, "healthy")
}
