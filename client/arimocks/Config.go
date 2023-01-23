// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *Config) Data(key *ari.Key) (*ari.ConfigData, error) {
	ret := _m.Called(key)

	var r0 *ari.ConfigData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ConfigData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ConfigData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: key
func (_m *Config) Delete(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *Config) Get(key *ari.Key) *ari.ConfigHandle {
	ret := _m.Called(key)

	var r0 *ari.ConfigHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ConfigHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ConfigHandle)
		}
	}

	return r0
}

// Update provides a mock function with given fields: key, tuples
func (_m *Config) Update(key *ari.Key, tuples []ari.ConfigTuple) error {
	ret := _m.Called(key, tuples)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, []ari.ConfigTuple) error); ok {
		r0 = rf(key, tuples)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConfig interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfig creates a new instance of Config. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfig(t mockConstructorTestingTNewConfig) *Config {
	mock := &Config{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
