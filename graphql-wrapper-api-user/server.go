package main

import (
	"fmt"
	"log"
	"net/http"
	"poc-graphql/graph"
	"poc-graphql/pkg/config"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("something went wrong during configuration loading %v", err)
	}

	initUserApiClient := graph.NewUserClient(*loadConfig)
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{UserClient: initUserApiClient}}))

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
	log.Printf("connect to http://localhost:%v/ for GraphQL playground", loadConfig.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(loadConfig.Server.Port), nil))
}
