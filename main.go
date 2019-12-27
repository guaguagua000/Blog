package main

import (
	"Blog/config"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	config.LoadBaseConfig()

	db, err := sql.Open("mysql", config.GetMySQLConfig())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
	router.Run(config.GetPort())
}
