package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"go.uber.org/zap"
)

type Server struct {
	router *gin.Engine
	logger *zap.Logger
}

func NewServer(logger *zap.Logger) (*Server, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	gso := GinServerOptions{
		BaseURL:     "",
		Middlewares: nil,
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"error": err.Error()})
		},
	}

	s := Server{
		r,
		logger,
	}

	spec, err := GetSwagger()
	if err != nil {
		return nil, err
	}

	r.Use(middleware.OapiRequestValidatorWithOptions(spec, &middleware.Options{
		SilenceServersWarning: true,
		ErrorHandler: func(c *gin.Context, err string, statusCode int) {
			c.JSON(statusCode, Message{Message: err})
		},
	}))
	RegisterHandlersWithOptions(r, &s, gso)

	return &s, nil
}

func (s *Server) Serve(bind string) error {
	return s.router.Run(bind)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Message: "pong"})
}
