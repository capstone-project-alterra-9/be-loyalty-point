// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// userController is an autogenerated mock type for the userController type
type userController struct {
	mock.Mock
}

type mockConstructorTestingTnewUserController interface {
	mock.TestingT
	Cleanup(func())
}

// newUserController creates a new instance of userController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newUserController(t mockConstructorTestingTnewUserController) *userController {
	mock := &userController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
