package utils

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		id    int64
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    args
		wantErr bool
	}{
		{
			name:    "generate token",
			args:    args{id: 1, email: "test@mail.com"},
			want:    args{id: 1, email: "test@mail.com"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenStr, err := GenerateToken(tt.args.id, tt.args.email)
			token, _ := getToken(tokenStr)
			claims := token.Claims.(jwt.MapClaims)
			got := args{id: int64(claims["userId"].(float64)), email: claims["email"].(string)}

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	validToken, _ := GenerateToken(1, "test@mail.com")
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "validate token",
			args:    args{tokenString: validToken},
			want:    1,
			wantErr: false,
		},
		{
			name:    "invalid token",
			args:    args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q3N0BtYWlsLmNvbSIsImV4cCI6MTczNjY3NzI3OSwidXNlcklkIjo2fQ.Qm0FbS0MACjf9Sv3pZ3tGn0sB2pqdYMGhnuhWl9-N24"},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
