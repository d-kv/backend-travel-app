// Code generated by mockery. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/d-kv/backend-travel-app/pkg/user-service/model"
)

// UserProvider is an autogenerated mock type for the UserProvider type
type UserProvider struct {
	mock.Mock
}

// Update provides a mock function with given fields: ctx, uuid, user
func (_m *UserProvider) Update(ctx context.Context, uuid string, user *model.User) error {
	ret := _m.Called(ctx, uuid, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.User) error); ok {
		r0 = rf(ctx, uuid, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// User provides a mock function with given fields: ctx, uuid
func (_m *UserProvider) User(ctx context.Context, uuid string) (*model.User, error) {
	ret := _m.Called(ctx, uuid)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.User, error)); ok {
		return rf(ctx, uuid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserProvider creates a new instance of UserProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserProvider(t mockConstructorTestingTNewUserProvider) *UserProvider {
	mock := &UserProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
