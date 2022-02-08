// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/burntcarrot/pm/entity/user"
	mock "github.com/stretchr/testify/mock"
)

// DomainRepo is an autogenerated mock type for the DomainRepo type
type DomainRepo struct {
	mock.Mock
}

// Register provides a mock function with given fields: ctx, domain
func (_m *DomainRepo) Register(ctx context.Context, domain user.Domain) (user.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, user.Domain) user.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, user.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *DomainRepo) GetByID(ctx context.Context, id string) (user.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) user.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password
func (_m *DomainRepo) Login(ctx context.Context, email string, password string) (user.Domain, error) {
	ret := _m.Called(ctx, email, password)

	var r0 user.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) user.Domain); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Get(0).(user.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
