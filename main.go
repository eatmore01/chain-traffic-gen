package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const URL_ENV_NAME = "URL"
const CONFIG_FILENAME = "env"
const CONFIG_TYPE = "env"
const RUNNING_PORT = "9000"

var url string

func init() {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w, using env varibales", err)

	}
	value_from_file := viper.GetString(URL_ENV_NAME)
	value_from_env := os.Getenv(URL_ENV_NAME)
	if value_from_env != "" {
		url = value_from_env
	}
	if value_from_env == "" && value_from_file != "" {
		url = value_from_file
	}

	if value_from_env == "" && value_from_file == "" {
		panic("Env url variables not set")
	}
}

func main() {
	fmt.Println(url)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "okay"})
	})

	router.POST("/", func(c *gin.Context) {
		resp, err := http.Get(url)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "External request failed",
			})
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{
				"error": "External service error",
				"body":  string(body),
			})
			return
		}

		c.Status(http.StatusOK)
	})

	port := fmt.Sprintf(": %s", RUNNING_PORT)

	router.Run(port)
}
