package divide

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"case1", 10, 2, 5, nil},
		{"case2", 10, 0, 0, errors.New("division by zero")},
		{"case3", 9, 3, 3, nil},
		{"case4", 7, -1, -7, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Fatalf("Divide() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Fatalf("Divide() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("Divide() = %d, want %d", got, tt.want)
			}
		})
	}
}
