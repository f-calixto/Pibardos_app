package user

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchQueryConstructor(t *testing.T) {
	// test 1 - valid email update
	req := []byte(`{"id":"12345678","email": "new_email"}`)
	r := UpsertUserRequest{}
	_ = json.Unmarshal(req, &r)
	query, args, err := PatchQueryConstructor(r)
	expectedQuery := "UPDATE users SET email = $1 WHERE id = $2"
	assert.Nil(t, err)
	assert.NotEmpty(t, args)
	assert.Equal(t, query, expectedQuery)
}

func ValidPatchQueryConstructor2(t *testing.T) {
	// test 2 - multiple valid updates
	req := []byte(`{"id":"12345678","email": "new_email", "status": "desayunando", "country": "br"}`)
	r := UpsertUserRequest{}
	_ = json.Unmarshal(req, &r)
	query, args, err := PatchQueryConstructor(r)
	expectedQuery := "UPDATE users SET email = $1, status = $2, country = $3 WHERE id = $4"
	assert.Nil(t, err)
	assert.NotEmpty(t, args)
	assert.Equal(t, query, expectedQuery)
}

func ValidPatchQueryConstructor4(t *testing.T) {
	// test 4 - multiple valid updates and one ignored field
	req := []byte(`{"id":"12345678","email": "new_email", "status": "desayunando", "country": "br", "created_at": "12/04/5666"}`)
	r := UpsertUserRequest{}
	_ = json.Unmarshal(req, &r)
	query, args, err := PatchQueryConstructor(r)
	expectedQuery := "UPDATE users SET email = $1, status = $2, country = $3 WHERE id = $4"
	assert.Nil(t, err)
	assert.NotEmpty(t, args)
	assert.Equal(t, query, expectedQuery)
}

func ValidPatchQueryConstructor3(t *testing.T) {
	// test 3 - no valid fields
	req := []byte(`{"id":"12345678"}`)
	r := UpsertUserRequest{}
	_ = json.Unmarshal(req, &r)
	query, args, err := PatchQueryConstructor(r)
	assert.Nil(t, args)
	assert.Equal(t, query, "")
	assert.Error(t, err)
}
