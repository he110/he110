package health_provider

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const logPrefix = "HEALTH_CHECK_SERVER"

type HealthServer struct {
	Port     string
	Endpoint string
	Logger   *zap.Logger
}

func NewHealthServer(port, endpoint string, l *zap.Logger) *HealthServer {
	return &HealthServer{
		Port:     port,
		Endpoint: endpoint,
		Logger:   l.With(zap.String("prefix", logPrefix)),
	}
}

func (s *HealthServer) ListenAndServe(ctx context.Context) error {
	router := mux.NewRouter()
	router.HandleFunc(s.Endpoint, healthHandler)

	server := http.Server{
		Addr:    ":" + s.Port,
		Handler: router,
	}

	errChan := make(chan error, 1)
	go func() {
		errChan <- server.ListenAndServe()
	}()
	s.Logger.Info("server started")

	select {
	case err := <-errChan:
		s.Logger.Warn("server down due to caught error")
		return err
	case <-ctx.Done():
		s.Logger.Info("server shut down by context")
		return server.Shutdown(ctx)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
