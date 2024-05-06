package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Event struct {
	ID          int
  Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      []int
}

var (
	events []Event = []Event{}
	id             = 0
)

func fetchEvents() []Event {
	return events
}

func getNextId() int {
	id += 1
	return id
}

func New(name string, description string, location string, datetime time.Time) (Event, error) {
	if name == "" || description == "" || location == "" {
		return Event{}, errors.New("Name ,desc, location cannot be empty")
	}

	event := Event{
		ID:          getNextId(),
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    datetime,
	}
	return event, nil
}

func (event Event) Save(){
  //later add to database
	events = append(events, event)
}
func GetAllEvents() string {
	events := fetchEvents()
	json, err := json.Marshal(events)
	fmt.Println(string(json))

	if err != nil {
		fmt.Println("Error while marshalling events to json: ", err)
		return ""
	}
	return string(json)
}
