package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mjeffers-998/interview-task/storage"
	"github.com/mjeffers-998/interview-task/structs"
)

type Controller struct {
	storage storage.Storage
}

func SetController(storage storage.Storage) {
	Cont = Controller{
		storage: storage,
	}
}

var Cont Controller

func CreateOrder(c *gin.Context) {
	//TODO: create a NewOrder() func with proper validation
	orderValue, err := strconv.ParseFloat(c.PostForm("order_value"), 32)
	if err != nil {
		c.AbortWithError(400, errors.New("order_value should be a float"))
	}
	storeID, err := strconv.Atoi(c.PostForm("store_id"))
	if err != nil {
		c.AbortWithError(400, errors.New("store_id should be an int"))
	}
	orderTime, err := time.Parse(time.RFC3339, c.PostForm("order_date"))
	if err != nil {
		c.AbortWithError(400, errors.New("order_date cannot be parsed"))
	}
	order := structs.Order{
		OrderDate:  orderTime,
		OrderValue: float32(orderValue),
		Status:     structs.Status(c.PostForm("status")),
		StoreID:    storeID,
	}
	id, err := Cont.storage.CreateOrder(&order)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"order_id": id,
	})
}
func CancelOrder(c *gin.Context) {
	//TODO
}

func GetRevenue(c *gin.Context) {
	//TODO
}

func ListAllOrders(c *gin.Context) {
	//TODO
}

func GetOrderByID(c *gin.Context) {
	//TODO
}
