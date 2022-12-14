package controllers

import (
	"GolangProject/dto"
	"GolangProject/helpers"
	"GolangProject/models"
	"GolangProject/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	All(context *gin.Context)
	Show(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService services.ProductService
	jwtService     services.JWTService
	//userService    services.UserService
}

func NewProductController(productServ services.ProductService, jwtServ services.JWTService) ProductController {
	return &productController{
		productService: productServ,
		jwtService:     jwtServ,
		//userService:    userServ,
	}
}

func (c *productController) All(context *gin.Context) {
	var products []models.Product = c.productService.All()
	res := helpers.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)
}

func (c *productController) Show(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product models.Product = c.productService.Show(id)
	if (product == models.Product{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", product)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) Insert(context *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		//authHeader := context.GetHeader("Authorization")
		//token, err := c.jwtService.ValidateToken(authHeader)
		//if err != nil {
		//	panic(err.Error())
		//}
		//claims := token.Claims.(jwt.MapClaims)
		//id := fmt.Sprintf("%v", claims["user_id"])
		//user := c.userService.Profile(id)
		result := c.productService.Insert(productCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *productController) Update(context *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	result := c.productService.Update(productUpdateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *productController) Delete(context *gin.Context) {
	var order models.Product
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	order.ID = id

	c.productService.Delete(order)
	res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
	context.JSON(http.StatusOK, res)
}

func (c *productController) getUserIDByToken(token string) interface{} {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
