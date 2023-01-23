// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// KeyOptionFunc is an autogenerated mock type for the KeyOptionFunc type
type KeyOptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *KeyOptionFunc) Execute(_a0 ari.Key) ari.Key {
	ret := _m.Called(_a0)

	var r0 ari.Key
	if rf, ok := ret.Get(0).(func(ari.Key) ari.Key); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(ari.Key)
	}

	return r0
}

type mockConstructorTestingTNewKeyOptionFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewKeyOptionFunc creates a new instance of KeyOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKeyOptionFunc(t mockConstructorTestingTNewKeyOptionFunc) *KeyOptionFunc {
	mock := &KeyOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
