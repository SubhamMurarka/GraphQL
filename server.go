package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	config "github.com/subhammurarka/GraphQL/Config"
	"github.com/subhammurarka/GraphQL/graph"
	repository "github.com/subhammurarka/GraphQL/repository"
	"github.com/subhammurarka/GraphQL/service"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	repo := repository.NewMaterialRepository()
	svc := service.NewMaterialService(repo)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{MaterialSvc: svc}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := config.Config.ServerPort

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
