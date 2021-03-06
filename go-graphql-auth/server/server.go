package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/prisma/prisma-examples/go-graphql-auth/prisma-client"
)

const defaultPort = "4000"

func main() {
	log.Printf("Initialize GraphQL service...")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := Resolver{
		Prisma: client,
	}
	r := chi.NewRouter()
	r.Use(AuthMiddleware)

	r.Handle("/", handler.Playground("GraphQL playground", "/query"))
	r.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &resolver})))

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
