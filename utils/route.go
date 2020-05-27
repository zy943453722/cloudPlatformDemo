package utils

import (
	"cloudPlatformDemo/middleware"
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerToFile())
	for _, opt := range options {
		opt(r)
	}
	return r
}
