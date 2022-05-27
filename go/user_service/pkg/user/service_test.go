package user

import (
	// std lib
	"testing"
	"time"

	// internal
	"github.com/coding-kiko/user_service/pkg/log"
	"github.com/stretchr/testify/assert"

	// third party
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// -------------------------------------------M O C K I N G ------------------------------------

type MockRepository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: id
func (_m *MockRepository) GetUser(id string) (User, error) {
	ret := _m.Called(id)

	var r0 User
	if rf, ok := ret.Get(0).(func(string) User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserGroups provides a mock function with given fields: id
func (_m *MockRepository) GetUserGroups(id string) ([]Group, error) {
	ret := _m.Called(id)

	var r0 []Group
	if rf, ok := ret.Get(0).(func(string) []Group); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: _a0
func (_m *MockRepository) InsertUser(_a0 User) (User, error) {
	ret := _m.Called(_a0)

	var r0 User
	if rf, ok := ret.Get(0).(func(User) User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAvatar provides a mock function with given fields: req
func (_m *MockRepository) UpdateAvatar(req UpdateAvatarRequest) (User, error) {
	ret := _m.Called(req)

	var r0 User
	if rf, ok := ret.Get(0).(func(UpdateAvatarRequest) User); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(UpdateAvatarRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: req
func (_m *MockRepository) UpdateUser(req UpsertUserRequest) (User, error) {
	ret := _m.Called(req)

	var r0 User
	if rf, ok := ret.Get(0).(func(UpsertUserRequest) User); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(UpsertUserRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type MockRabbitProducer struct {
	mock.Mock
}

// AvatarQueue provides a mock function with given fields: file
func (_m *MockRabbitProducer) AvatarQueue(file File) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(File) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// -------------------------------------------------------------------------------------------

var sampleUser1 = User{
	Id:         uuid.NewString(),
	Created_at: time.Now().UTC().String(),
	Email:      "kikoaudi2001@gmail.com",
	Username:   "kiko_audi",
	Country:    "BR",
	Birthdate:  "17/04/2001",
	Status:     "",
	Avatar:     defaultUserAvatar,
}

var (
	uruguay   = "UY"
	newStatus = "desayunando"
)

func TestUpsertUser(t *testing.T) {
	logger := log.NewLogger()

	testCases := []struct {
		testName      string
		request       UpsertUserRequest
		response      func(req UpsertUserRequest) (User, error)
		checkResponse func(t *testing.T, user User, err error)
	}{
		{
			testName: "valid: insert new user",
			request: UpsertUserRequest{
				Id:         &sampleUser1.Id,
				Email:      &sampleUser1.Email,
				Created_at: &sampleUser1.Created_at,
				Username:   &sampleUser1.Username,
				Country:    &sampleUser1.Country,
				Status:     &sampleUser1.Status,
				Birthdate:  &sampleUser1.Birthdate,
			},
			response: func(req UpsertUserRequest) (User, error) {
				return sampleUser1, nil
			},
			checkResponse: func(t *testing.T, user User, err error) {
				assert.Nil(t, err)

			},
		},
		{
			testName: "valid: update user",
			request: UpsertUserRequest{
				Id:      &sampleUser1.Id,
				Country: &uruguay,
				Status:  &newStatus,
			},
			response: func(req UpsertUserRequest) (User, error) {
				cpy := sampleUser1
				cpy.Country = uruguay
				cpy.Status = newStatus
				return cpy, nil
			},
			checkResponse: func(t *testing.T, user User, err error) {
				assert.Nil(t, err)

			},
		},
	}

	// initialize srevice layer passing mocked dependencies
	mockRepo := new(MockRepository)
	mockRepo.AssertExpectations(t)
	mockRabbit := new(MockRabbitProducer)
	service := NewService(mockRepo, mockRabbit, logger)

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			repoRes, err := tc.response(tc.request)
			mockRepo.On("UpdateUser", tc.request).Return(repoRes, err)
			mockRepo.On("InsertUser", tc.request).Return(repoRes, err)
			res, err := service.UpsertUser(tc.request)
			tc.checkResponse(t, res, err)
		})
	}
}
