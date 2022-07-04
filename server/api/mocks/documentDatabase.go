// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/fabian-lapotre/document-api/server/model"
	mock "github.com/stretchr/testify/mock"
)

// documentDatabase is an autogenerated mock type for the documentDatabase type
type documentDatabase struct {
	mock.Mock
}

// CreateDocument provides a mock function with given fields: doc
func (_m *documentDatabase) CreateDocument(doc model.Document) error {
	ret := _m.Called(doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Document) error); ok {
		r0 = rf(doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDocumentByID provides a mock function with given fields: id
func (_m *documentDatabase) DeleteDocumentByID(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDocumentByID provides a mock function with given fields: id
func (_m *documentDatabase) GetDocumentByID(id uint) (model.Document, error) {
	ret := _m.Called(id)

	var r0 model.Document
	if rf, ok := ret.Get(0).(func(uint) model.Document); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Document)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDocuments provides a mock function with given fields:
func (_m *documentDatabase) GetDocuments() ([]model.Document, error) {
	ret := _m.Called()

	var r0 []model.Document
	if rf, ok := ret.Get(0).(func() []model.Document); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewDocumentDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// newDocumentDatabase creates a new instance of documentDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDocumentDatabase(t mockConstructorTestingTnewDocumentDatabase) *documentDatabase {
	mock := &documentDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
