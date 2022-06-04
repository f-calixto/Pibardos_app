package calendar

import (
	// std lib
	"testing"

	// third party
	"github.com/stretchr/testify/assert"
)

func TestCheckDateCollision(t *testing.T) {
	testCases := []struct {
		name          string
		s1            string
		e1            string
		s2            string
		e2            string
		checkResponse func(t *testing.T, err error)
	}{
		{
			name: "valid new event",
			s1:   "2022-06-02T16:00:00.000Z",
			e1:   "2022-06-02T20:00:00.000Z",
			s2:   "2022-06-02T22:00:00.000Z",
			e2:   "2022-06-02T23:00:00.000Z",
			checkResponse: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "valid new event",
			s1:   "2022-06-02T16:00:00.000Z",
			e1:   "2022-06-02T20:00:00.000Z",
			s2:   "2022-06-02T20:00:00.000Z",
			e2:   "2022-06-02T23:00:00.000Z",
			checkResponse: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name: "invalid new event",
			s1:   "2022-06-02T16:00:00.000Z",
			e1:   "2022-06-02T20:00:00.000Z",
			s2:   "2022-06-02T18:00:00.000Z",
			e2:   "2022-06-02T23:00:00.000Z",
			checkResponse: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name: "valid new event",
			s1:   "2022-06-02T16:00:00.000Z",
			e1:   "2022-06-02T20:00:00.000Z",
			s2:   "2022-06-02T12:00:00.000Z",
			e2:   "2022-06-02T17:00:00.000Z",
			checkResponse: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckDateCollision(tc.s1, tc.e1, tc.s2, tc.e2)
			tc.checkResponse(t, err)
		})
	}
}
