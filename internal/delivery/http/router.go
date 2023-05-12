package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo/internal/usecase"
	"todo/pkg/logger"
)

func NewRouter(handler *gin.Engine, log logger.Interface, uc usecase.UseCases) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h := handler.Group("/v1")
	{
		newTaskRoutes(h, log, uc.Task)
	}
}
