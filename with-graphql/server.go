package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Arnobkumarsaha/learn-nats/with-graphql/graph"
	"github.com/Arnobkumarsaha/learn-nats/with-graphql/graph/generated"
)

// https://dev.to/karanpratapsingh/graphql-subscriptions-at-scale-with-nats-f19

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	cfg := generated.Config{
		Resolvers: &graph.Resolver{
			Nats: nc,
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
