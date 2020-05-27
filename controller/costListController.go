package controller

import (
	"cloudPlatformDemo/middleware"
	"cloudPlatformDemo/models"
	"strconv"

	//"fmt"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	//"os"
)

type ResultJson struct {
	List 		[]models.Cost 	`json:"list"`
	TotalCount 	int64 			`json:"totalCount"`
}

func GetCostList(c *gin.Context) {
	productLine, err := strconv.ParseInt(c.Query("productLine"), 10, 64)
	if err != nil {
		middleware.Log.Error("string convert int fail", err)
		c.JSON(200, gin.H{
			"status": 10005,
			"message": err.Error(),
		})
		return
	}
	var productType int64
	productType, err = strconv.ParseInt(c.Query("productType"), 10, 64)
	if err != nil {
		middleware.Log.Error("string convert int fail", err)
		c.JSON(200, gin.H{
			"status": 10005,
			"message": err.Error(),
		})
		return
	}
	validator := validator2.New()
	costValidator := models.CostValidator{
		ProductLine: productLine,
		ProductType: productType,
	}
	if c.Query("ipType") != "" {
		costValidator.IpType = c.Query("ipType")
	}
	middleware.Log.Info("costValidator is:", costValidator)
	if err = validator.Struct(costValidator); err != nil {
		middleware.Log.Error("validator fail ", err)
		c.JSON(200, gin.H{
			"status": 10001,
			"message": err.Error(),
		})
		return
	}
	var result []models.Cost
	var count int64
	result, count, err = models.GetCostList(&costValidator)
	if err != nil {
		middleware.Log.Error("query db fail", err)
		c.JSON(200, gin.H{
			"status": 10002,
			"message": err.Error(),
		})
	} else {
		resultJson := ResultJson{
			List: result,
			TotalCount: count,
		}
		c.JSON(200, gin.H{
			"status": 10000,
			"message": "success",
			"data": resultJson,
		})
	}
}
