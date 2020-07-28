package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	l *zap.Logger
)

func init() {
	newer := zap.NewProduction
	zapLogger, err := newer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l = zapLogger
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		l.Info(
			"API is called",
			zap.String("path", c.FullPath()),
			zap.String("handler", c.HandlerName()),
		)
		c.JSON(http.StatusOK, "ok!")
	})
	router.Run(":3000")
}
