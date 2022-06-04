package calendar

import (
	// std lib
	"testing"

	// third party
	"github.com/stretchr/testify/assert"
)

func TestAddOneHour(t *testing.T) {
	testCases := []struct {
		name            string
		startDate       string
		expectedEndDate string
		checkResponse   func(t *testing.T, startDate string, expectedQuery string)
	}{
		{
			name:            "add one hour",
			startDate:       "2022-06-02T22:00:00Z",
			expectedEndDate: "2022-06-02T23:00:00Z",
			checkResponse: func(t *testing.T, result string, expectedEndDate string) {
				assert.Equal(t, result, expectedEndDate)
			},
		},
		{
			name:            "add one hour and start new day",
			startDate:       "2022-06-02T23:30:00Z",
			expectedEndDate: "2022-06-03T00:30:00Z",
			checkResponse: func(t *testing.T, result string, expectedEndDate string) {
				assert.Equal(t, result, expectedEndDate)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			endDate := AddOneHour(tc.startDate)
			tc.checkResponse(t, endDate, tc.expectedEndDate)
		})
	}
}

func TestCheckEndDateValidity(t *testing.T) {
	testCases := []struct {
		name          string
		startDate     string
		endDate       string
		checkResponse func(t *testing.T, err error)
	}{
		{
			name:      "valid end date",
			startDate: "2022-06-02T22:00:00Z",
			endDate:   "2022-06-02T23:00:00Z",
			checkResponse: func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		{
			name:      "invalid end date - too early",
			startDate: "2022-06-02T22:00:00Z",
			endDate:   "2022-06-02T21:00:00Z",
			checkResponse: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		{
			name:      "invalid end date - equal to start date",
			startDate: "2022-06-02T22:00:00Z",
			endDate:   "2022-06-02T22:00:00Z",
			checkResponse: func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := CheckEndDateValidity(tc.endDate, tc.startDate)
			tc.checkResponse(t, err)
		})
	}
}
