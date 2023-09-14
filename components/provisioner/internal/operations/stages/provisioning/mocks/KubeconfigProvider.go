// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// KubeconfigProvider is an autogenerated mock type for the KubeconfigProvider type
type KubeconfigProvider struct {
	mock.Mock
}

// FetchFromShoot provides a mock function with given fields: shootName
func (_m *KubeconfigProvider) FetchFromShoot(shootName string) ([]byte, error) {
	ret := _m.Called(shootName)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(shootName)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(shootName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(shootName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKubeconfigProvider creates a new instance of KubeconfigProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKubeconfigProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *KubeconfigProvider {
	mock := &KubeconfigProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
