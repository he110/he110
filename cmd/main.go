package main

import (
	"log"
	"net/http"

	"He110/PersonalWebSite/internal/graph"
	"He110/PersonalWebSite/internal/graph/resolvers"
	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file was found")
	}
}

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatal("cannot initialize app config")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/playground", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
