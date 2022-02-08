// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	project "github.com/burntcarrot/pm/entity/project"
	mock "github.com/stretchr/testify/mock"
)

// DomainRepo is an autogenerated mock type for the DomainRepo type
type DomainRepo struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: ctx, domain
func (_m *DomainRepo) CreateProject(ctx context.Context, domain project.Domain) (project.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 project.Domain
	if rf, ok := ret.Get(0).(func(context.Context, project.Domain) project.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(project.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, project.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectByID provides a mock function with given fields: ctx, username, projectID
func (_m *DomainRepo) GetProjectByID(ctx context.Context, username string, projectID string) ([]project.Domain, error) {
	ret := _m.Called(ctx, username, projectID)

	var r0 []project.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []project.Domain); ok {
		r0 = rf(ctx, username, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]project.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, username, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjects provides a mock function with given fields: ctx, username
func (_m *DomainRepo) GetProjects(ctx context.Context, username string) ([]project.Domain, error) {
	ret := _m.Called(ctx, username)

	var r0 []project.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) []project.Domain); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]project.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
