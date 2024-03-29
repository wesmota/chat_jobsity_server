// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "github.com/wesmota/go-jobsity-chat-server/models"
)

// ChatRoomRepo is an autogenerated mock type for the ChatRoomRepo type
type ChatRoomRepo struct {
	mock.Mock
}

// CreateChatMessage provides a mock function with given fields: ctx, chatMessage
func (_m *ChatRoomRepo) CreateChatMessage(ctx context.Context, chatMessage models.ChatMessage) error {
	ret := _m.Called(ctx, chatMessage)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatMessage) error); ok {
		r0 = rf(ctx, chatMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateChatRoom provides a mock function with given fields: ctx, chatRoom
func (_m *ChatRoomRepo) CreateChatRoom(ctx context.Context, chatRoom models.ChatRoom) error {
	ret := _m.Called(ctx, chatRoom)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ChatRoom) error); ok {
		r0 = rf(ctx, chatRoom)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *ChatRoomRepo) CreateUser(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListChatRooms provides a mock function with given fields: ctx
func (_m *ChatRoomRepo) ListChatRooms(ctx context.Context) ([]models.ChatRoom, error) {
	ret := _m.Called(ctx)

	var r0 []models.ChatRoom
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]models.ChatRoom, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []models.ChatRoom); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ChatRoom)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatRoomRepo creates a new instance of ChatRoomRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatRoomRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatRoomRepo {
	mock := &ChatRoomRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
