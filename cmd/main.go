package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"He110/PersonalWebSite/internal/providers"
	"He110/PersonalWebSite/internal/providers/gql_provider"
	"He110/PersonalWebSite/internal/providers/health_provider"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	errStopped = errors.New("service stopped")
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file was found")
	}
}

func main() {
	t1, _ := zap.NewProduction()
	container, err := providers.BuildContainer()
	if err != nil {
		t1.Fatal("cannot initialize dependencies", zap.Error(err))
	}

	gr, ctx := errgroup.WithContext(context.Background())

	err = container.Invoke(func(
		gqlServer *gql_provider.GqlServer,
		healthServer *health_provider.HealthServer,
	) {
		gr.Go(func() error {
			return gqlServer.ListenAndServe(ctx)
		})
		gr.Go(func() error {
			return healthServer.ListenAndServe(ctx)
		})
		gr.Go(func() error {
			signals := make(chan os.Signal, 1)
			signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
			defer signal.Stop(signals)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-signals:
				return errStopped
			}
		})
		if err := gr.Wait(); err != nil && err != errStopped {
			t1.Fatal("terminating due to caught error", zap.Error(err))
		}
	})

	if err != nil {
		t1.Fatal("dependency invoke failed", zap.Error(err))
	}
}
