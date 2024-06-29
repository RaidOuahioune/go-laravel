package graphql

import (
	"fmt"

	"demo.com/hello/core/graphql/graph"
	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/gin-gonic/gin"
)

func GraphQLHandler() gin.HandlerFunc {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		fmt.Println("GraphQLHandler")
		srv.ServeHTTP(c.Writer, c.Request)
	}

}
