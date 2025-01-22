package services

import (
	"strconv"
	"testing"
	"time"

	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/common/models"
)

func TestRegisterForEvent(t *testing.T) {
	db.InitDB("sqlite3", "../../../test.db")
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
			name: "register for event",
			args: args{
				eventId: strconv.FormatInt(ev.ID, 10),
				userId:  u.ID,
			},
			wantErr: false,
		},
		{
			name: "register for event fail",
			args: args{
				eventId: "999999",
				userId:  99999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterForEvent(tt.args.eventId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("RegisterForEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCancelRegistration(t *testing.T) {
	db.InitDB("sqlite3", "../../../test.db")
	u := models.NewUser("test@test.com", "testpass")
	u.Create()

	ev := models.NewEvent("test event", "test event description", "test location", time.Now().UTC())
	ev.UserID = u.ID
	ev.Create()

	RegisterForEvent(strconv.FormatInt(ev.ID, 10), u.ID)

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
			name: "cancel registration",
			args: args{
				eventId: strconv.FormatInt(ev.ID, 10),
				userId:  u.ID,
			},
			wantErr: false,
		},
		{
			name: "cancel registration fail",
			args: args{
				eventId: "999999",
				userId:  99999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CancelRegistration(tt.args.eventId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("CancelRegistration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
