package main

import (
	"fmt"
	//"log"
	"github.com/gin-gonic/gin"
	//"backendAPI/models"
	//"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	//"context"
	//"backendAPI/controllers"
	"github.com/jinzhu/gorm"
)
var db *gorm.DB
var err error
func main() {
	r := gin.Default()
	dbUri := "<POSTGRESQL_URI>"
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
		})
	r.POST("/companyinfo", AddNewCompany)
	r.Run() 
}

type Company struct {
	Name string `json:"name"`
	Tel string `json:"tel"`
	Contactperson string `json:"contactperson"`
}

func AddNewCompany (c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
// Validate input
    var input Company
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
// Create Book
    company := Company{Name: input.Name, Tel: input.Tel, Contactperson: input.Contactperson}
    db.Create(&company)
    c.JSON(http.StatusOK, gin.H{"data": company})
}
