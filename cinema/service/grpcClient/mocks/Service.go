// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	authen "cinema/service/grpcClient/protobuf/auth"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// VerifyToken provides a mock function with given fields: data
func (_m *Service) VerifyToken(data *authen.TokenRequest) (*authen.TokenReply, error) {
	ret := _m.Called(data)

	var r0 *authen.TokenReply
	if rf, ok := ret.Get(0).(func(*authen.TokenRequest) *authen.TokenReply); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*authen.TokenReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*authen.TokenRequest) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
