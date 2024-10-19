package database

import (
	"testing"
)

func TestNewSQL(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(t *testing.T)
		wantDB bool
	}{
		{
			name: "success",
			setup: func(t *testing.T) {
				t.Setenv("SQL_USERNAME", "testing")
				t.Setenv("SQL_PASSWORD", "testing")
				t.Setenv("SQL_HOST", "testing")
				t.Setenv("SQL_PORT", "testing")
				t.Setenv("SQL_DATABASE_NAME", "testing")
			},
			wantDB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			got := NewSQL()

			if (got != nil) != tt.wantDB {
				t.Errorf("NewSQL() = %v, wantDB %v", got, tt.wantDB)
			}
		})
	}
}
