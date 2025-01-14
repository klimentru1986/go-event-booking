package services

import (
	"errors"
	"strconv"

	"github.com/klimentru1986/go-event-booking/common/models"
)

func CreateEvent(event *models.Event, userId int64) error {

	event.UserID = userId
	err := event.Create()

	if err != nil {
		return err
	}

	return nil
}

func UpdateEvent(eventId string, updatedEvent *models.Event, userId int64) error {
	id, event, err := FindEventByStrId(eventId)

	if err != nil {
		return err
	}

	if userId != event.UserID {
		return errors.New("Unauthorized")
	}

	updatedEvent.ID = *id

	err = updatedEvent.Update()

	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(eventId string, userId int64) error {
	_, event, err := FindEventByStrId(eventId)

	if err != nil {
		return err
	}

	if userId != event.UserID {
		return errors.New("Unauthorized")
	}

	err = event.Delete()

	if err != nil {
		return err
	}

	return nil
}

func FindEventByStrId(strId string) (*int64, *models.Event, error) {
	id, err := strconv.ParseInt(strId, 10, 64)

	if err != nil {
		return nil, nil, err
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		return nil, nil, err
	}

	return &id, event, nil
}
