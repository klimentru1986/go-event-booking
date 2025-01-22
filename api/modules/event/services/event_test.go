package services

import (
	"strconv"
	"testing"
	"time"

	"github.com/klimentru1986/go-event-booking/common/config"
	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/common/models"
)

var conf_test = config.New("../../../test.env")

func TestFindEventByStrId(t *testing.T) {
	db.InitDB(conf_test.DbDriver, conf_test.DbSource)
	ev := models.NewEvent(
		"Event1",
		"Descr1",
		"Location1",
		time.Now().UTC(),
	)
	ev.Create()

	type args struct {
		strId string
	}
	tests := []struct {
		name    string
		args    args
		want    *int64
		want1   *models.Event
		wantErr bool
	}{
		{
			name: "find event by str id",
			args: args{
				strId: strconv.FormatInt(ev.ID, 10),
			},
			want:    &ev.ID,
			want1:   &ev,
			wantErr: false,
		},
		{
			name: "fail find event by str id",
			args: args{
				strId: "9999",
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := FindEventByStrId(tt.args.strId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindEventByStrId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && *got != *tt.want {
				t.Errorf("FindEventByStrId() got = %v, want %v", got, tt.want)
			}
			if got != nil && tt.want != nil && (got1.Name != tt.want1.Name || got1.Description != tt.want1.Description) {
				t.Errorf("FindEventByStrId() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeleteEvent(t *testing.T) {
	db.InitDB(conf_test.DbDriver, conf_test.DbSource)
	u := models.NewUser("test@test.com", "testpass")
	u.Create()

	ev := models.NewEvent("test event", "test event description", "test location", time.Now().UTC())
	ev.UserID = u.ID
	ev.Create()

	type args struct {
		eventId string
		userId  int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete event by id",
			args: args{
				eventId: strconv.FormatInt(ev.ID, 10),
				userId:  ev.UserID,
			},
			wantErr: false,
		},
		{
			name: "delete event by id with wrong user",
			args: args{
				eventId: strconv.FormatInt(ev.ID, 10),
				userId:  0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteEvent(tt.args.eventId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateEvent(t *testing.T) {
	db.InitDB(conf_test.DbDriver, conf_test.DbSource)
	u := models.NewUser("test@test.com", "testpass")
	u.Create()

	ev := models.NewEvent("test event", "test event description", "test location", time.Now().UTC())

	t.Run("create event", func(t *testing.T) {
		err := CreateEvent(&ev, u.ID)
		if err != nil {
			t.Errorf("CreateEvent() error = %v", err)
		}
	})

}
