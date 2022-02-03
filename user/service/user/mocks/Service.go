// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	input "user/service/user/input"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	output "user/service/user/output"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// All provides a mock function with given fields: ctx
func (_m *Service) All(ctx context.Context) ([]*output.User, error) {
	ret := _m.Called(ctx)

	var r0 []*output.User
	if rf, ok := ret.Get(0).(func(context.Context) []*output.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*output.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, in
func (_m *Service) Create(ctx context.Context, in *input.UserInput) (*output.User, *string, error) {
	ret := _m.Called(ctx, in)

	var r0 *output.User
	if rf, ok := ret.Get(0).(func(context.Context, *input.UserInput) *output.User); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.User)
		}
	}

	var r1 *string
	if rf, ok := ret.Get(1).(func(context.Context, *input.UserInput) *string); ok {
		r1 = rf(ctx, in)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *input.UserInput) error); ok {
		r2 = rf(ctx, in)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: ctx, in
func (_m *Service) Delete(ctx context.Context, in *input.UserInput) error {
	ret := _m.Called(ctx, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *input.UserInput) error); ok {
		r0 = rf(ctx, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: ctx, in
func (_m *Service) Read(ctx context.Context, in *input.UserInput) (*output.User, error) {
	ret := _m.Called(ctx, in)

	var r0 *output.User
	if rf, ok := ret.Get(0).(func(context.Context, *input.UserInput) *output.User); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *input.UserInput) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, in
func (_m *Service) Update(ctx context.Context, in *input.UserInput) (*output.User, error) {
	ret := _m.Called(ctx, in)

	var r0 *output.User
	if rf, ok := ret.Get(0).(func(context.Context, *input.UserInput) *output.User); ok {
		r0 = rf(ctx, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *input.UserInput) error); ok {
		r1 = rf(ctx, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, in, filename, file
func (_m *Service) Upload(ctx context.Context, in *input.UserInput, filename string, file multipart.File) (*output.User, error) {
	ret := _m.Called(ctx, in, filename, file)

	var r0 *output.User
	if rf, ok := ret.Get(0).(func(context.Context, *input.UserInput, string, multipart.File) *output.User); ok {
		r0 = rf(ctx, in, filename, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *input.UserInput, string, multipart.File) error); ok {
		r1 = rf(ctx, in, filename, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
