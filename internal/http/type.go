package http

import (
	"github.com/gin-gonic/gin"
)

// Server encapsulates the complexity of serving http.  It includes a
// router and a hooks system.
type Server struct {
	*gin.Engine
}
