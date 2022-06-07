package user

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchQueryConstructor(t *testing.T) {
	testCases := []struct {
		name          string
		reqJson       string
		expectedQuery string
		checkResponse func(t *testing.T, expectedQuery string, query string, args []interface{}, err error)
	}{
		{
			name:          "valid email update",
			reqJson:       `{"id":"12345678","email": "new_email"}`,
			expectedQuery: "UPDATE users SET email = $1 WHERE id = $2",
			checkResponse: func(t *testing.T, expectedQuery string, query string, args []interface{}, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, args)
				assert.Equal(t, query, expectedQuery)
			},
		},
		{
			name:          "multiple valid updates",
			reqJson:       `{"id":"12345678","email": "new_email", "status": "desayunando", "country": "br"}`,
			expectedQuery: "UPDATE users SET email = $1, country = $2, status = $3 WHERE id = $4",
			checkResponse: func(t *testing.T, expectedQuery string, query string, args []interface{}, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, args)
				assert.Equal(t, query, expectedQuery)
			},
		},
		{
			name:          "multiple valid updates and one ignored field",
			reqJson:       `{"id":"12345678","email": "new_email", "status": "desayunando", "country": "br", "created_at": "12/04/5666"}`,
			expectedQuery: "UPDATE users SET email = $1, country = $2, status = $3 WHERE id = $4",
			checkResponse: func(t *testing.T, expectedQuery string, query string, args []interface{}, err error) {
				assert.Nil(t, err)
				assert.NotEmpty(t, args)
				assert.Equal(t, query, expectedQuery)
			},
		},
		{
			name:          "no valid fields",
			reqJson:       `{"id":"12345678"}`,
			expectedQuery: "",
			checkResponse: func(t *testing.T, expectedQuery string, query string, args []interface{}, err error) {
				assert.Nil(t, args)
				assert.Equal(t, query, expectedQuery)
				assert.Error(t, err)
			},
		},
		{
			name:          "empty request",
			reqJson:       `{}`,
			expectedQuery: "",
			checkResponse: func(t *testing.T, expectedQuery string, query string, args []interface{}, err error) {
				assert.Nil(t, args)
				assert.Equal(t, query, expectedQuery)
				assert.Error(t, err)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := UpsertUserRequest{}
			_ = json.Unmarshal([]byte(tc.reqJson), &r)
			query, args, err := PatchQueryConstructor(r)
			tc.checkResponse(t, tc.expectedQuery, query, args, err)
		})
	}
}
