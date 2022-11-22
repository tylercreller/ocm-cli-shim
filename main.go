package main

import (
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type ocmData struct {
	Token string `json:"token"`
	URL   string `json:"url"`
}

func main() {
	router := gin.Default()
	router.GET("/", getOCM)

	router.Run("0.0.0.0:8082")
}

func getOCM(c *gin.Context) {
	token, _ := exec.Command("ocm", "token").Output()
	apiUrl, _ := exec.Command("ocm", "config", "get", "url").Output()
	c.IndentedJSON(http.StatusOK, ocmData{
		Token: strings.TrimSpace(string(token)),
		URL:   strings.TrimSpace(string(apiUrl)),
	})
}
