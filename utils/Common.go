package utils

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.elastic.co/apm/module/apmchi"
)

//JSONMarshal is func
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func health(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"status": "ok",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func liveness(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"status": "UP",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

//NewRouterPlus router with health check
func NewRouterPlus() *chi.Mux {
	router := chi.NewRouter()

	router.Use(apmchi.Middleware())
	router.Get("/health", health)
	router.Get("/liveness", liveness)
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
	// disable pprof  sue to security issue , date  4 mei 2020
	// router.Mount("/debug", mid.Profiler())

	return router
}
