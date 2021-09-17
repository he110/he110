package gql_provider

import (
	"context"
	"net/http"

	"He110/PersonalWebSite/internal/graph"
	"He110/PersonalWebSite/internal/graph/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	logPrefix       = "GQL_SERVER"
	playgroundTitle = "GraphQL playground"
)

type GqlServer struct {
	Port               string
	MainEndpoint       string
	PlaygroundEndpoint string
	Logger             *zap.Logger
	Resolver           *resolvers.Resolver
}

func NewGqlServer(Resolver *resolvers.Resolver, Port, MainEndpoint, PlaygroundEndpoint string, l *zap.Logger) *GqlServer {
	return &GqlServer{
		Resolver:           Resolver,
		Port:               Port,
		MainEndpoint:       MainEndpoint,
		PlaygroundEndpoint: PlaygroundEndpoint,
		Logger:             l.With(zap.String("prefix", logPrefix)),
	}
}

func (s *GqlServer) ListenAndServe(ctx context.Context) error {
	cfg := graph.Config{Resolvers: s.Resolver}
	schema := graph.NewExecutableSchema(cfg)
	srv := handler.NewDefaultServer(schema)

	errChan := make(chan error, 1)
	server := http.Server{
		Addr:    ":" + s.Port,
		Handler: s.prepareHandler(srv),
	}
	go func() {
		errChan <- server.ListenAndServe()
	}()

	s.Logger.Info("gql server is ready to be served on port " + s.Port)

	select {
	case err := <-errChan:
		s.Logger.Warn("gql stopped because of error")
		return err
	case <-ctx.Done():
		err := server.Shutdown(ctx)
		s.Logger.Info("gql server was shut downed by context")
		return err
	}
}

func (s *GqlServer) prepareHandler(srv *handler.Server) http.Handler {
	router := mux.NewRouter()
	router.Handle(s.PlaygroundEndpoint, playground.Handler(playgroundTitle, s.MainEndpoint))
	router.Handle(s.MainEndpoint, srv)

	return router
}
