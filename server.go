package main

import (
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/Akshaytermin/gqlbeego/graph"
	"github.com/Akshaytermin/gqlbeego/graph/generated"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@/beego?charset=utf8")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)

	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	h := handler.Playground("GraphQL", "/query")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)

	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	orm.Debug = true

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
