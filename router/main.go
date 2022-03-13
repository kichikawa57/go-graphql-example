package router

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/kichikawa/auth"
	"github.com/kichikawa/ent"
	"github.com/kichikawa/graph/generated"
	"github.com/kichikawa/graph/resolver"
)

func graphqlHandler() gin.HandlerFunc {

	fmt.Println("tets")

	return func(c *gin.Context) {
		client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			"db", "5432", "postgres", "development", "password"))

		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}

		h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
			Client: client,
		}}))

		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, tokenErr := auth.GetHeaderToken(c)

		if tokenErr != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": tokenErr.Error(),
			})
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(AuthMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	return r
}
