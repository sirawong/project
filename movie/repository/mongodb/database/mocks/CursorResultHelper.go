// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CursorResultHelper is an autogenerated mock type for the CursorResultHelper type
type CursorResultHelper struct {
	mock.Mock
}

// Close provides a mock function with given fields: ctx
func (_m *CursorResultHelper) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Decode provides a mock function with given fields: v
func (_m *CursorResultHelper) Decode(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields: ctx
func (_m *CursorResultHelper) Next(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}