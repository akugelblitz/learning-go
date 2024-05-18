package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events try again later"},
		)
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventsById(context *gin.Context) {
	strId := context.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println("Could not convert string to int in id")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invaild format of id"})
		return
	}
	event, err := models.GetEventsById(id)
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events try again later"},
		)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": id, "event": event})
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println("Binding error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = event.Save()
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not create events try again later"},
		)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func updateEvent(context *gin.Context) {
	strId := context.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println("Could not convert string to int in id")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invaild format of id"})
		return
	}
	_, err = models.GetEventsById(id)
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events try again later"},
		)
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		fmt.Println("Binding error", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = int64(id)
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated"})
}

func deleteEvent(context *gin.Context) {
	strId := context.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		fmt.Println("Could not convert string to int in id")
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invaild format of id"})
		return
	}
	event, err := models.GetEventsById(id)
	if err != nil {
		fmt.Println(err)
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not fetch events try again later"},
		)
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted"})
}
