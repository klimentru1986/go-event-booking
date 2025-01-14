package services

func RegisterForEvent(eventId string, userId int64) error {

	_, event, err := FindEventByStrId(eventId)

	if err != nil {
		return err
	}

	err = event.RegisterUser(userId)

	if err != nil {
		return err
	}

	return nil
}

func CancelRegistration(eventId string, userId int64) error {
	_, event, err := FindEventByStrId(eventId)

	if err != nil {

		return err
	}

	err = event.CancelRegistration(userId)

	if err != nil {
		return err
	}

	return nil
}
