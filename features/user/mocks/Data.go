// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	user "SistemVaksinAPI/features/user"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// EditPasswordByID provides a mock function with given fields: data, newPassword
func (_m *Data) EditPasswordByID(data user.UserCore, newPassword string) (user.UserCore, error) {
	ret := _m.Called(data, newPassword)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore, string) user.UserCore); ok {
		r0 = rf(data, newPassword)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore, string) error); ok {
		r1 = rf(data, newPassword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EditUser provides a mock function with given fields: data
func (_m *Data) EditUser(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: data
func (_m *Data) InsertUser(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: data
func (_m *Data) Login(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserByID provides a mock function with given fields: id
func (_m *Data) SelectUserByID(id int) (user.UserCore, error) {
	ret := _m.Called(id)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(int) user.UserCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectUserEmail provides a mock function with given fields: data
func (_m *Data) SelectUserEmail(data user.UserCore) (user.UserCore, error) {
	ret := _m.Called(data)

	var r0 user.UserCore
	if rf, ok := ret.Get(0).(func(user.UserCore) user.UserCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(user.UserCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.UserCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
