package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Butternut01/inventory-service/internal/entity"
    "github.com/Butternut01/inventory-service/internal/usecase"
)

type ProductController struct {
    productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) *ProductController {
    return &ProductController{
        productUseCase: productUseCase,
    }
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
    var product entity.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.productUseCase.CreateProduct(&product); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
    id := ctx.Param("id")

    product, err := c.productUseCase.GetProduct(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
    id := ctx.Param("id")

    var product entity.Product
    if err := ctx.ShouldBindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product.ID = id
    if err := c.productUseCase.UpdateProduct(&product); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
    id := ctx.Param("id")

    if err := c.productUseCase.DeleteProduct(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusNoContent, nil)
}

func (c *ProductController) ListProducts(ctx *gin.Context) {
    var filter entity.ProductFilter

    // Parse query parameters
    if name := ctx.Query("name"); name != "" {
        filter.Name = name
    }
    if category := ctx.Query("category"); category != "" {
        filter.Category = category
    }
    if minPrice := ctx.Query("min_price"); minPrice != "" {
        if val, err := strconv.ParseFloat(minPrice, 64); err == nil {
            filter.MinPrice = val
        }
    }
    if maxPrice := ctx.Query("max_price"); maxPrice != "" {
        if val, err := strconv.ParseFloat(maxPrice, 64); err == nil {
            filter.MaxPrice = val
        }
    }
    if page := ctx.Query("page"); page != "" {
        if val, err := strconv.Atoi(page); err == nil && val > 0 {
            filter.Page = val
        } else {
            filter.Page = 1
        }
    } else {
        filter.Page = 1
    }
    if limit := ctx.Query("limit"); limit != "" {
        if val, err := strconv.Atoi(limit); err == nil && val > 0 {
            filter.Limit = val
        } else {
            filter.Limit = 10
        }
    } else {
        filter.Limit = 10
    }

    products, err := c.productUseCase.ListProducts(filter)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, products)
}