package product

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/response"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

func setupControllers(engine *gin.Engine) {
	r := engine.Group("/product")
	r.GET("", GetAllProducts)
	r.POST("/add", AddProduct)
	r.POST("/addMany", AddManyProduct)
}

// AddProduct adds a single product.
// @Summary Add Product
// @Description Add a single product
// @Produce json
// @Success 201 {object} string "created"
// @Failure 400 {object} string "error"
// @Router /product/add [post]
func AddProduct(c *gin.Context) {
	var payload Product

	err := c.BindJSON(&payload)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	err = Add(c, &payload)
	if err != nil {
		logger.Error("Unable to Save to Product Error:", err)
		response.BadRequest(c, err)
		return
	}
	response.SuccessfullyCreate(c, "Product Created")
}

// AddManyProduct adds many products at once.
// @Summary Add Many Product
// @Description Add Many product at once
// @Produce json
// @Success 201 {object} string "created"
// @Failure 400 {object} string "error"
// @Router /product/addMany [post]
func AddManyProduct(c *gin.Context) {
	var payload []Product

	err := c.BindJSON(&payload)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	err = AddMany(c, &payload)
	if err != nil {
		logger.Error("Unable to Save to Product Error:", err)
		response.BadRequest(c, err)
		return
	}
	response.SuccessfullyCreate(c, "Product Created")
}

// GetAllProducts retrieves all products.
// @Summary Get all Products
// @Description Get all products
// @Produce json
// @Success 200 {array} string "ok"
// @Failure 400 {object} string "error"
// @Router /product/ [get]
func GetAllProducts(c *gin.Context) {
	data, err := Fetch(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	response.Success(c, data)
}
