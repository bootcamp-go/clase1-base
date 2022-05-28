package handler

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/clase1-base/internal/models"
	internal "github.com/bootcamp-go/clase1-base/internal/products"
	"github.com/bootcamp-go/clase1-base/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Service internal.Service
}

func NewHandlerProducts(s internal.Service) *Product {
	return &Product{Service: s}
}

func (p *Product) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.String(400, "Invalid id")
		}
		product, err := p.Service.GetOne(int(id))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("There was an error: %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, product, ""))
		}
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var prod models.Product

		err := ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.String(400, "there was an error %v", err)
		} else {
			response, err := p.Service.Store(prod)
			if err != nil {
				ctx.String(400, "there was an error %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var prod models.Product

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.String(400, "Invalid id")
		}

		err = ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.String(400, "there was an error %v", err)
		} else {
			updatedProduct, err := p.Service.Update(int(id), prod)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, updatedProduct)
			}
		}

	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "Invalid id")
		}

		err = p.Service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "The product %d has been removed", id)
		}

	}
}
