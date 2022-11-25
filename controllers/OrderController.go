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

type OrderController interface {
	All(context *gin.Context)
	Show(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type orderController struct {
	orderService services.OrderService
	jwtService   services.JWTService
}

func NewOrderController(orderServ services.OrderService, jwtServ services.JWTService) OrderController {
	return &orderController{
		orderService: orderServ,
		jwtService:   jwtServ,
	}
}

func (c *orderController) All(context *gin.Context) {
	var orders []models.Order = c.orderService.All()
	res := helpers.BuildResponse(true, "OK", orders)
	context.JSON(http.StatusOK, res)
}

func (c *orderController) Show(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var order models.Order = c.orderService.Show(id)
	if (order == models.Order{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", order)
		context.JSON(http.StatusOK, res)
	}
}

func (c *orderController) Insert(context *gin.Context) {
	var orderCreateDTO dto.OrderCreateDTO
	errDTO := context.ShouldBind(&orderCreateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			orderCreateDTO.UserID = convertedUserID
		}
		result := c.orderService.Create(orderCreateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *orderController) Update(context *gin.Context) {
	var orderUpdateDTO dto.OrderUpdateDTO
	errDTO := context.ShouldBind(&orderUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.orderService.IsAllowedToEdit(userID, orderUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			orderUpdateDTO.UserID = id
		}
		result := c.orderService.Update(orderUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *orderController) Delete(context *gin.Context) {
	var order models.Order
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	order.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.orderService.IsAllowedToEdit(userID, order.ID) {
		c.orderService.Delete(order)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *orderController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
