package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	//"backendAPI/models"
	"database/sql"
	_ "github.com/lib/pq"
	"context"
)
var db *sql.DB
var err error
func main() {
	r := gin.Default()
	connString := "<POSTGRESQL_URI>"
	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	log.Printf("Connected!\n")
	defer db.Close()
	CheckVersion()
	r.POST("/companyinfo", AddNewCompany)
	r.POST("/testJSON", exampleJSON)
	r.Run() 
}
func CheckVersion() {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Ping database failed:", err.Error())
	}

	var version string
	err = db.QueryRowContext(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	log.Printf("%s\n", version)
}
func AddNewCompany(c *gin.Context) {
    var input Company
	e := c.ShouldBindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	log.Printf(input.name)
	c.JSON(200, gin.H{
		"name": input.name,
		"tel": input.tel,
		"contactperson": input.contactperson,
	})
}
type Company struct {
	name string `json:"name"`
	tel string `json:"tel"`
	contactperson string `json:"contactperson"`
}

func exampleJSON(c *gin.Context) {
	var input Result
	e := c.ShouldBindJSON(&input)
	if e != nil {
		fmt.Println(e)
	}
	log.Printf(input.Name)
	c.JSON(200, gin.H{
		"id": input.ID,
		"name": input.Name,
		"message": input.Message,
	})
}

type Result struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Message string `json:"message"`
}