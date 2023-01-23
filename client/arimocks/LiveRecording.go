// Code generated by mockery v2.16.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v6"
	mock "github.com/stretchr/testify/mock"
)

// LiveRecording is an autogenerated mock type for the LiveRecording type
type LiveRecording struct {
	mock.Mock
}

// Data provides a mock function with given fields: key
func (_m *LiveRecording) Data(key *ari.Key) (*ari.LiveRecordingData, error) {
	ret := _m.Called(key)

	var r0 *ari.LiveRecordingData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.LiveRecordingData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LiveRecordingData)
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

// Get provides a mock function with given fields: key
func (_m *LiveRecording) Get(key *ari.Key) *ari.LiveRecordingHandle {
	ret := _m.Called(key)

	var r0 *ari.LiveRecordingHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.LiveRecordingHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.LiveRecordingHandle)
		}
	}

	return r0
}

// Mute provides a mock function with given fields: key
func (_m *LiveRecording) Mute(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pause provides a mock function with given fields: key
func (_m *LiveRecording) Pause(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Resume provides a mock function with given fields: key
func (_m *LiveRecording) Resume(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scrap provides a mock function with given fields: key
func (_m *LiveRecording) Scrap(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields: key
func (_m *LiveRecording) Stop(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stored provides a mock function with given fields: key
func (_m *LiveRecording) Stored(key *ari.Key) *ari.StoredRecordingHandle {
	ret := _m.Called(key)

	var r0 *ari.StoredRecordingHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.StoredRecordingHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.StoredRecordingHandle)
		}
	}

	return r0
}

// Subscribe provides a mock function with given fields: key, n
func (_m *LiveRecording) Subscribe(key *ari.Key, n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(*ari.Key, ...string) ari.Subscription); ok {
		r0 = rf(key, n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}

// Unmute provides a mock function with given fields: key
func (_m *LiveRecording) Unmute(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLiveRecording interface {
	mock.TestingT
	Cleanup(func())
}

// NewLiveRecording creates a new instance of LiveRecording. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLiveRecording(t mockConstructorTestingTNewLiveRecording) *LiveRecording {
	mock := &LiveRecording{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
