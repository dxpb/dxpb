package http

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("http.bind", "")
	viper.SetDefault("http.port", 8080)
}

// New returns  an http.Server that is  bound and ready to  serve, but
// isn't serving yet.
func New() (Server, error) {
	gin.SetMode(gin.ReleaseMode)

	x := Server{
		gin.New(),
	}

	return x, nil
}

// RegisterGinHook takes in a path to register at, and a function to
// handle that path.
func (s *Server) RegisterGinHook(path, verb string, f gin.HandlerFunc) {
	s.Handle(verb, path, f)
}

// Serve blocks forever and runs the webserver.
func (s *Server) Serve() {
	addr := viper.GetString("http.bind")
	port := viper.GetInt("http.port")
	log.Panic(s.Run(fmt.Sprintf("%s:%d", addr, port)))
}
