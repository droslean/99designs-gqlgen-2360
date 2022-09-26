package main

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/droslean/99designs-gqlgen-2360/graph"
	"github.com/droslean/99designs-gqlgen-2360/graph/generated"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {

	router := chi.NewRouter()

	config := generated.Config{Resolvers: graph.NewResolver()}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	logrus.Info("Listening on 8080")
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":8080"), router))
}
