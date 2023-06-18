package storage

import (
	"os"
	"testing"
)

func TestNewDB(t *testing.T) {
	tests := []struct {
		name    string
		want    *DB
		wantErr bool
	}{
		{"test.db", nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewDB()
			defer os.Remove("./" + tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
