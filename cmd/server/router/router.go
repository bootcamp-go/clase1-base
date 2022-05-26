package router

import (
	"database/sql"

	"github.com/bootcamp-go/clase1-base/cmd/server/handler"
	internal "github.com/bootcamp-go/clase1-base/internal/products"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r, r.Group("/api/v1"), db}
}

func (r *router) MapRoutes() {
	r.buildProductsRoutes()
}

func (r *router) buildProductsRoutes() {
	repo := internal.NewRepository(r.db)
	service := internal.NewService(repo)
	handler := handler.NewHandlerProducts(service)

	r.rg.Group("/products")
	{
		r.rg.GET("", handler.GetOne())
		r.rg.POST("", handler.Store())
		r.rg.PUT("/:id", handler.Update())
		r.rg.DELETE("/:id", handler.Delete())
	}
}
