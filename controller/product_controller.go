package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
    productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
        productUseCase: usecase,
    }
}

func (p *productController) GetProducts(ctx *gin.Context) {
	
    products, err := p.productUseCase.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
    }

    ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

    var product model.Product
    err := ctx.ShouldBindJSON(&product)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err)
        return
    }


    insertedProduct, err := p.productUseCase.CreateProduct(product)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    ctx.JSON(http.StatusCreated, insertedProduct)

}

func (p *productController) GetProductsById(ctx *gin.Context) {
    id := ctx.Param("product_id")
    if id == "" {
        response := model.Response{ Message: "Product ID is required" }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    productId, err := strconv.Atoi(id)
    if err != nil {
        response := model.Response{ Message: "Product ID must be a number" }
        ctx.JSON(http.StatusBadRequest, response)
        return
    }

    product, err := p.productUseCase.GetProductById(productId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }
    
    if product == nil {
        response := model.Response{ Message: "Product not found" }
        ctx.JSON(http.StatusNotFound, response)
        return
    }

    ctx.JSON(http.StatusOK, product)
}

