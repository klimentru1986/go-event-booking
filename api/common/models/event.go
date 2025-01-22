package models

import (
	"fmt"
	"time"

	"github.com/klimentru1986/go-event-booking/common/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required" `
	Description string    `json:"description" binding:"required" `
	Location    string    `json:"location" binding:"required" `
	DateTime    time.Time `json:"date_time" binding:"required" `
	UserID      int64     `json:"user_id"`
}

func NewEvent(name string, description string, location string, dateTime time.Time) Event {
	return Event{
		Name:        name,
		Description: description,
		Location:    location,
		DateTime:    dateTime,
	}
}

func (e *Event) Create() error {
	query := fmt.Sprintf(`INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES ('%s', '%s', '%s', '%s', '%x') RETURNING id`,
		e.Name, e.Description, e.Location, e.DateTime.Format(time.RFC3339), e.UserID)

	var id int64
	err := db.DB.QueryRow(query).Scan(&id)
	if err != nil {
		return err
	}

	e.ID = id

	return err
}

func (e *Event) Update() error {
	query := `
		UPDATE events 
		SET name = $1, description = $2, location = $3, dateTime = $4
		WHERE id = $5
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	query := "DELETE FROM events WHERE id = $1"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e *Event) RegisterUser(userId int64) error {
	query := `
		INSERT INTO registrations(user_id, event_id)
		VALUES ($1, $2)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, e.ID)

	return err
}

func (e *Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE user_id = $1 AND event_id = $2
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, e.ID)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events = []Event{}

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = $1"
	row := db.DB.QueryRow(query, id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}
	return &e, nil
}
