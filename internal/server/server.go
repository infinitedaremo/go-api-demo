package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infinitedaremo/go-api-demo/internal/app"
	middleware "github.com/oapi-codegen/gin-middleware"
	"go.uber.org/zap"
)

type Server struct {
	router        *gin.Engine
	logger        *zap.Logger
	personService app.PersonService
}

func NewServer(logger *zap.Logger, personService app.PersonService) (*Server, error) {
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
		personService,
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

func (s *Server) GetPortfolio(c *gin.Context, id PersonID) {
	portfolio, err := s.personService.GetPortfolio(c, int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, Message{Message: "pong"})
}
