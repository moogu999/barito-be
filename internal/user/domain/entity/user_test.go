package entity

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	email := "testing@testing.com"
	password := "testing"

	tests := []struct {
		name     string
		email    string
		password string
		want     User
		wantErr  bool
	}{
		{
			name:     "success",
			email:    email,
			password: password,
			want: User{
				Email:     email,
				CreatedBy: email,
			},
			wantErr: false,
		},
		{
			name:     "password is too long",
			email:    email,
			password: "7DA1LRHz7KsRCS0dvO5A1CvjE5jDXDh2Z9iPeN1741260y8a9K2ze738aJxOztz7TRQ8lBLdZ",
			want:     User{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := NewUser(tt.email, tt.password)

			if tt.wantErr {
				if err == nil {
					t.Error("NewUser() want err but got nil")
				}
			} else {
				if tt.want.ID != got.ID {
					t.Errorf("NewUser().ID = %d, want %d", got.ID, tt.want.ID)
				}
				if tt.want.Email != got.Email {
					t.Errorf("NewUser().Email = %s, want %s", got.Email, tt.want.Email)
				}
				if tt.want.CreatedBy != got.CreatedBy {
					t.Errorf("NewUser().CreatedBy = %s, want %s", got.CreatedBy, tt.want.CreatedBy)
				}
				if got.CreatedAt.IsZero() {
					t.Error("NewUser().CreatedAt.IsZero() == true")
				}
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	t.Parallel()

	defaultPwd := "testing"
	user, _ := NewUser("testing@testing.com", defaultPwd)
	tests := []struct {
		user     User
		password string
		name     string
		want     bool
	}{
		{
			user:     user,
			password: defaultPwd,
			name:     "success",
			want:     true,
		},
		{
			user:     user,
			password: "incorrect",
			name:     "incorrect password",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := tt.user.VerifyPassword(tt.password)

			if tt.want != got {
				t.Errorf("VerifyPassword() = %t, want %t", got, tt.want)
			}
		})
	}
}
