package services

import (
	"strconv"

	"github.com/klimentru1986/go-event-booking/common/models"
)

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
