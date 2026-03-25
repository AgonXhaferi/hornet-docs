package main

import (
	"hornet-docs/internal/application"
	"hornet-docs/internal/infrastructure/adapter"
	"hornet-docs/internal/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := adapter.NewInMemoryRepository()

	documentService := application.NewDocumentService(repo)

	handler := http.NewDocumentHandler(documentService)

	r := gin.Default()

	r.POST("/documents", handler.Create)
	r.GET("/documents/:id", handler.Get)
	r.DELETE("/documents/:id", handler.Delete)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
