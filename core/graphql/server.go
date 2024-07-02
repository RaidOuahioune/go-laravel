package graphql

import (
	"demo.com/hello/core/graphql/graph"
	"demo.com/hello/core/http/auth"
	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/gin-gonic/gin"
)

func GraphQLHandler() gin.HandlerFunc {

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {

		var user = auth.CurrentUser(c)
		c.Set("user", user)
		srv.ServeHTTP(c.Writer, c.Request.WithContext(c))
	}

}
