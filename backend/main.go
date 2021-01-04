package main

import (
	//"fmt"
	//"log"
	"backendAPI/models"

	"github.com/gin-gonic/gin"

	//"database/sql"
	_ "github.com/lib/pq"
	//"context"

	"backendAPI/controllers"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var db *gorm.DB
var NoSQL *sqlx.DB
var err error

func main() {
	r := gin.Default()
	db := models.ConnectPsql()
	NoSQL := models.ConnectPNoSQL()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.Use(func(c *gin.Context) {
		c.Set("NoSQL", NoSQL)
		c.Next()
	})
	//r.POST("/AddCompany", controllers.AddNewCompany)
	//r.POST("/AddAccount", controllers.AddNewAccount)
	r.GET("/testData/:id", controllers.GetTestData)
	r.GET("/getHost/:scan_id", controllers.GetHost)
	r.GET("/getVal/:scan_id", controllers.GetVul)
	r.POST("/addEndpoint", controllers.AddNewEndpoint)
	r.POST("/addScan", controllers.AddNewScan)
	r.POST("/RAW", controllers.Raw)
	r.PATCH("/updateScan/:status", controllers.UpdateScan)
	r.GET("/getEndpoints/:google_id", controllers.GetEndpoints)
	r.DELETE("/DeleteEndpoint/:endpoint", controllers.DeleteEndpoint)
	r.Run()
}

