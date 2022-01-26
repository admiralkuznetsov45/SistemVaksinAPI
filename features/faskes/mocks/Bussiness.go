// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	faskes "SistemVaksinAPI/features/faskes"

	mock "github.com/stretchr/testify/mock"
)

// Bussiness is an autogenerated mock type for the Bussiness type
type Bussiness struct {
	mock.Mock
}

// CreateFaskes provides a mock function with given fields: data
func (_m *Bussiness) CreateFaskes(data faskes.FaskesCore) (faskes.FaskesCore, error) {
	ret := _m.Called(data)

	var r0 faskes.FaskesCore
	if rf, ok := ret.Get(0).(func(faskes.FaskesCore) faskes.FaskesCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(faskes.FaskesCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(faskes.FaskesCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllFaskes provides a mock function with given fields:
func (_m *Bussiness) GetAllFaskes() []faskes.FaskesCore {
	ret := _m.Called()

	var r0 []faskes.FaskesCore
	if rf, ok := ret.Get(0).(func() []faskes.FaskesCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]faskes.FaskesCore)
		}
	}

	return r0
}

// GetFaskesByID provides a mock function with given fields: id
func (_m *Bussiness) GetFaskesByID(id int) (faskes.FaskesCore, error) {
	ret := _m.Called(id)

	var r0 faskes.FaskesCore
	if rf, ok := ret.Get(0).(func(int) faskes.FaskesCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(faskes.FaskesCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFaskesByName provides a mock function with given fields: data
func (_m *Bussiness) GetFaskesByName(data faskes.FaskesCore) ([]faskes.FaskesCore, error) {
	ret := _m.Called(data)

	var r0 []faskes.FaskesCore
	if rf, ok := ret.Get(0).(func(faskes.FaskesCore) []faskes.FaskesCore); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]faskes.FaskesCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(faskes.FaskesCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
