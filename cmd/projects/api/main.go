package main

import (
	"example.com/api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func populateDummyEvents(num int) {
	var err error
	var event models.Event
	for range num {
		event, err = models.New("test event", "test desc", "whatup", time.Now())
		event.Save()
	}
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	server := gin.Default()
	populateDummyEvents(10)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	}

	event.ID = 1
	event.UserId = []int{1}

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}
