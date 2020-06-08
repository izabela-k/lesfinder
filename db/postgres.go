package db

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)


func Postgres() gin.HandlerFunc {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Panicf("Unable to connection to database\n")
	}

	return func(c *gin.Context) {
		c.Set("Postgres", conn)
		c.Next()
	}
}