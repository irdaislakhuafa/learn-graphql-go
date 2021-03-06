package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/irdaislakhuafa/learn-graphql-go/graph"
	"github.com/irdaislakhuafa/learn-graphql-go/graph/generated"
	"github.com/irdaislakhuafa/learn-graphql-go/src/database"
	"github.com/irdaislakhuafa/learn-graphql-go/src/repositories"
	"github.com/irdaislakhuafa/learn-graphql-go/src/schema"
)

const defaultPort = "8080"

func main() {

	// connect to database
	dbcon := database.DBConnection{
		Driver:   "mysql",
		Username: "root",
		Password: "password",
		DBName:   "learn_graphql_go",
		DBPort:   "3306",
	}
	dbcon.Connect()

	// generate table schema
	schemaGenerator := database.SchemaGenerator{
		Connection: dbcon.GetConnection(),
	}
	schemaGenerator.GenerateSchema(schema.Users, schema.Todo)

	// define repositories
	userRepository := &repositories.UserRepository{
		Connection: dbcon.GetConnection(),
	}
	todoRepository := &repositories.TodoRepository{
		Connection: dbcon.GetConnection(),
	}

	// run graphql
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserRepository: userRepository,
		TodoRepository: todoRepository,
		Connection:     dbcon.GetConnection(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
