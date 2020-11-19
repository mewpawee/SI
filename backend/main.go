package main

import (
	//"fmt"
	//"log"
	"github.com/gin-gonic/gin"
	"backendAPI/models"
	//"database/sql"
	_ "github.com/lib/pq"
	//"context"
	"backendAPI/controllers"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB
var err error
func main() {
	r := gin.Default()
	db := models.ConnectPsql()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
		})
	//r.POST("/AddCompany", controllers.AddNewCompany)
	//r.POST("/AddAccount", controllers.AddNewAccount)
	r.GET("/testData/:id", controllers.GetTestData)
	r.GET("/getHost/:scan_id", controllers.GetHost)
	r.GET("/getVal/:scan_id", controllers.GetVul)
	//r.POST("/AddLog", controllers.AddLog)
	r.Run() 
}
