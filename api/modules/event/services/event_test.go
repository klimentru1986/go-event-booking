package services

import (
	"strconv"
	"testing"
	"time"

	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/common/models"
)

func TestFindEventByStrId(t *testing.T) {
	db.InitDB("../../../test.db")
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
