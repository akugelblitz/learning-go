package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"example.com/api/models"
)

func populateDummyEvents(num int) {
	var err error
	for range num {
		_, err = models.New("test event", "test desc", "whatup", time.Now())
	}
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := gin.Default()
	populateDummyEvents(10)
	server.GET("/events", getEvents)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}
