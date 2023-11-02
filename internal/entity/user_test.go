package entity

import (
	"testing"
)

var (
	name     = "Test"
	email    = "test@test.com"
	password = "123456"
)

func TestUser_ValidatePassword(t *testing.T) {
	type fields struct {
		Name     string
		Email    string
		Password string
	}
	type args struct {
		password string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should validate password",
			fields: fields{
				Name:     name,
				Email:    email,
				Password: password,
			},
			args: args{
				password: password,
			},
			want: true,
		},
		{
			name: "Should not validate password",
			fields: fields{
				Name:     name,
				Email:    email,
				Password: password,
			},
			args: args{
				password: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := NewUser(tt.fields.Name, tt.fields.Email, tt.fields.Password)

			if got := u.ValidatePassword(tt.args.password); got != tt.want {
				t.Errorf("User.ValidatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
