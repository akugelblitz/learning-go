package models

import (
	"fmt"
	"time"

	"example.com/api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

var events []Event = []Event{}

func fetchEvents() []Event {
	return events
}

func GetAllEvents() ([]Event, error) {
	query := `
  SELECT * FROM events
  `
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserId,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventsById(id int) (*Event, error) {
	query := fmt.Sprintf("SELECT * FROM events WHERE id=?")
	rows := db.DB.QueryRow(query, id)

	var event Event
	err := rows.Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&event.DateTime,
		&event.UserId,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Save() error {
	query := `
  INSERT INTO events(name, desc, location, dateTime, user_id)
  VALUES (?, ?, ?, ?, ?)
  `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime,
		event.UserId,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id

	return err
}

func (event Event) Update() error {
	query := `
  UPDATE events
  SET name = ?, desc = ?, location = ?, datetime = ?
  WHERE id = ?
  `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `
  DELETE FROM events
  WHERE id = ?
  `

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
