package services

import (
	"reflect"
	"testing"

	"github.com/klimentru1986/go-event-booking/common/db"
	"github.com/klimentru1986/go-event-booking/common/dto"
	"github.com/klimentru1986/go-event-booking/common/models"
)

func TestSignup(t *testing.T) {
	testUserEmail := "testuser@test.com"
	testUserPassword := "testpassword777"
	userDto := dto.CreateUserDto{Email: testUserEmail, Password: testUserPassword}

	type args struct {
		userDto *dto.CreateUserDto
	}
	type want struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "valid user",
			args: args{
				userDto: &userDto,
			},
			want: want{
				email:    testUserEmail,
				password: testUserPassword,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		db.InitDB("sqlite3", "../../../test.db")
		u := models.NewUser(tt.args.userDto.Email, tt.args.userDto.Password)
		u.Delete()

		t.Run(tt.name, func(t *testing.T) {
			newUser, err := Signup(tt.args.userDto)
			got := want{
				email:    newUser.Email,
				password: newUser.Password,
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Signup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	testUserEmail := "testuser@test.com"
	testUserPassword := "testpassword777"
	userDto := dto.CreateUserDto{Email: testUserEmail, Password: testUserPassword}

	invalidUserDto := dto.CreateUserDto{Email: testUserEmail, Password: "invalidPass"}

	type args struct {
		userDto *dto.CreateUserDto
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "login",
			args: args{
				userDto: &userDto,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "login false",
			args: args{
				userDto: &invalidUserDto,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		db.InitDB("sqlite3", "../../../test.db")
		u := models.NewUser(userDto.Email, userDto.Password)
		u.Delete()
		Signup(&userDto)

		t.Run(tt.name, func(t *testing.T) {
			token, err := Login(tt.args.userDto)
			got := token != ""
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
