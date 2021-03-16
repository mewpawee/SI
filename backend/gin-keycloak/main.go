package main

import (
	"flag"
	"time"

	"github.com/gin-gonic/gin"
	ginglog "github.com/szuecs/gin-glog"
	"github.com/tbaehler/gin-keycloak/pkg/ginkeycloak"
)

func main() {
	flag.Parse()
	router := gin.New()
	router.Use(ginglog.Logger(3 * time.Second))
	router.Use(ginkeycloak.RequestLogger([]string{"uid"}, "data"))
	router.Use(gin.Recovery())

	var sbbEndpoint = ginkeycloak.KeycloakConfig{
		Url:   "https://csi.cmkl.ac.th",
		Realm: "csi",
	}

	config := ginkeycloak.BuilderConfig{
		Service: "test",
		Url:     "https://csi.cmkl.ac.th",
		Realm:   "csi",
	}
	privateGroup := router.Group("/api/privateGroup")
	privateGroup.Use(ginkeycloak.Auth(ginkeycloak.AuthCheck(), sbbEndpoint))
	privateGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
	privateUser := router.Group("/api/privateUser")
	privateUser.Use(ginkeycloak.NewAccessBuilder(config).
		RestrictButForRealm("broker_firm").
		Build())
	privateUser.GET("/", func(c *gin.Context) {
		uid, okUID := c.Get("uid")
		if okUID {
			c.JSON(200, gin.H{"message": uid})
		}
	})
	router.Run()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "ping"})
	// })

	// r.Run()
}
