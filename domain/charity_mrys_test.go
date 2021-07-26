package domain

import (
	"errors"
	"testing"
)

var (
	ErrCharityMrysNotFound = errors.New("charity mrys not found")
)

func TestCharityMrys(t *testing.T) {
	t.Parallel()

	// type args struct {
	// 	name string
	// }

	tests := []struct {
		name        string
		charityMrys CharityMrys
		args        CharityMrys
		expected    CharityMrys
		expectedErr error
	}{
		{
			name:        "Success creating charity mrys",
			charityMrys: CharityMrys{},
			args: CharityMrys{
				Name: "Test",
			},
			expected: CharityMrys{
				Name: "Test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.charityMrys.Name = tt.args.Name

			if tt.charityMrys != tt.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tt.name,
					tt.charityMrys,
					tt.expected,
				)
			}
		})
	}

}
