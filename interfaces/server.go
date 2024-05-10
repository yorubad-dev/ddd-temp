package interfaces

import (
	"log/slog"

	"github.com/KingDaemonX/ddd-template/domain/repository/interfaces/middlewares"
	"github.com/gin-gonic/gin"
)

type server struct {
	Router *gin.Engine
	mw     *middlewares.Middleware
	ph     *Project
}

func NewServer(slg *slog.Logger, mw *middlewares.Middleware, ph *Project) *server {
	return &server{
		Router: gin.Default(),
		mw:     mw,
		ph:     ph,
	}
}
