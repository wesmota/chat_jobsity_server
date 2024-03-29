// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/wesmota/go-jobsity-chat-server/models"
)

// AuthorizationRepo is an autogenerated mock type for the AuthorizationRepo type
type AuthorizationRepo struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *AuthorizationRepo) CreateUser(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *AuthorizationRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ret := _m.Called(ctx, email)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthorizationRepo creates a new instance of AuthorizationRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthorizationRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthorizationRepo {
	mock := &AuthorizationRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
