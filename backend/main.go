package main

import (
	//"fmt"
	//"log"
	"backendAPI/models"
	"flag"
	"time"

	"github.com/gin-gonic/gin"

	//"database/sql"
	_ "github.com/lib/pq"
	//"context"

	"backendAPI/controllers"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"

	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)
var db *gorm.DB
var NoSQL *sqlx.DB
var err error

func main() {
	flag.Parse()
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
	//prawee code
	r.Use(ginglog.Logger(3 * time.Second))
	r.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	r.Use(gin.Recovery())

	var sbbEndpoint = ginkeycloak.KeycloakConfig{
		Url:   "https://csi.cmkl.ac.th",
		Realm: "csi",
	}

	config := ginkeycloak.BuilderConfig{
		Service: "test",
		Url:     "https://csi.cmkl.ac.th",
		Realm:   "csi",
	}
	privateGroup := r.Group("/api/privateGroup")
	privateGroup.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), sbbEndpoint))
	privateGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
	privateUser := r.Group("/api/privateUser")
	privateUser.Use(ginkeycloak.NewAccessBuilder(config).
		RestrictButForRealm("broker_firm").
		Build())
	/*privateUser.GET("/", func(c *gin.Context) {
		uid, okUID := c.Get("uid")
		if okUID {
			c.JSON(200, gin.H{"message": uid})
		}
	})*/
	privateUser.GET("/addPool", controllers.AddPool)
	privateUser.GET("/addEndpoint", controllers.AddEndpoint)
	privateUser.GET("/deleteEndpoint", controllers.DeleteEndpoint)
	//privateUser.GET("/addEndpoint", controllers.AddEndpoint)
	privateUser.GET("/getPools", controllers.GetCompanyPools)
	privateUser.GET("/getEndpoints", controllers.GetPoolEndpoints)
	//prawee code end
	//r.POST("/AddCompany", controllers.AddNewCompany)
	//r.POST("/AddAccount", controllers.AddNewAccount)
	/*r.GET("/testData/:id", controllers.GetTestData)
	r.GET("/getHost/:scan_id", controllers.GetHost)
	r.GET("/getVal/:scan_id", controllers.GetVul)
	r.POST("/addEndpoint", controllers.AddNewEndpoint)
	r.POST("/addScan", controllers.AddNewScan)*/
	r.POST("/resultLogs", controllers.Raw)
	/*r.PATCH("/updateScan/:status", controllers.UpdateScan)
	r.GET("/getEndpoints/:google_id", controllers.GetEndpoints)
	r.DELETE("/DeleteEndpoint/:endpoint", controllers.DeleteEndpoint)*/
	r.Run()
}

