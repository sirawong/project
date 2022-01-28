// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "user/entities"

	input "user/service/auth/input"

	mock "github.com/stretchr/testify/mock"

	output "user/service/auth/output"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GenAuth provides a mock function with given fields: ctx, ent
func (_m *Service) GenAuth(ctx context.Context, ent *entities.User) (*string, error) {
	ret := _m.Called(ctx, ent)

	var r0 *string
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) *string); ok {
		r0 = rf(ctx, ent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.User) error); ok {
		r1 = rf(ctx, ent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in
func (_m *Service) Login(ctx context.Context, in *input.AuthInput) (*output.User, *string, error) {
	ret := _m.Called(ctx, in)

	var r0 *output.User
	if rf, ok := ret.Get(0).(func(context.Context, *input.AuthInput) *output.User); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.User)
		}
	}

	var r1 *string
	if rf, ok := ret.Get(1).(func(context.Context, *input.AuthInput) *string); ok {
		r1 = rf(ctx, in)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *input.AuthInput) error); ok {
		r2 = rf(ctx, in)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Logout provides a mock function with given fields: ctx, in
func (_m *Service) Logout(ctx context.Context, in *input.AuthInput) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *input.AuthInput) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LogoutAll provides a mock function with given fields: ctx, in
func (_m *Service) LogoutAll(ctx context.Context, in *input.AuthInput) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *input.AuthInput) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: ctx, encodedToken
func (_m *Service) Verify(ctx context.Context, encodedToken string) (string, *output.User, error) {
	ret := _m.Called(ctx, encodedToken)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, encodedToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *output.User
	if rf, ok := ret.Get(1).(func(context.Context, string) *output.User); ok {
		r1 = rf(ctx, encodedToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*output.User)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, encodedToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}