// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gregoryusip/first-project/controller (interfaces: ProductControllerModel)

// Package mocks is a generated GoMock package.
package mocks

import (
	json "encoding/json"
	jrpc2 "github.com/bitwurx/jrpc2"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProductControllerModel is a mock of ProductControllerModel interface
type MockProductControllerModel struct {
	ctrl     *gomock.Controller
	recorder *MockProductControllerModelMockRecorder
}

// MockProductControllerModelMockRecorder is the mock recorder for MockProductControllerModel
type MockProductControllerModelMockRecorder struct {
	mock *MockProductControllerModel
}

// NewMockProductControllerModel creates a new mock instance
func NewMockProductControllerModel(ctrl *gomock.Controller) *MockProductControllerModel {
	mock := &MockProductControllerModel{ctrl: ctrl}
	mock.recorder = &MockProductControllerModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductControllerModel) EXPECT() *MockProductControllerModelMockRecorder {
	return m.recorder
}

// AddProduct mocks base method
func (m *MockProductControllerModel) AddProduct(arg0 json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*jrpc2.ErrorObject)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct
func (mr *MockProductControllerModelMockRecorder) AddProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductControllerModel)(nil).AddProduct), arg0)
}

// DeletedProduct mocks base method
func (m *MockProductControllerModel) DeletedProduct(arg0 json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletedProduct", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*jrpc2.ErrorObject)
	return ret0, ret1
}

// DeletedProduct indicates an expected call of DeletedProduct
func (mr *MockProductControllerModelMockRecorder) DeletedProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletedProduct", reflect.TypeOf((*MockProductControllerModel)(nil).DeletedProduct), arg0)
}

// ReadedProduct mocks base method
func (m *MockProductControllerModel) ReadedProduct(arg0 json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadedProduct", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*jrpc2.ErrorObject)
	return ret0, ret1
}

// ReadedProduct indicates an expected call of ReadedProduct
func (mr *MockProductControllerModelMockRecorder) ReadedProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadedProduct", reflect.TypeOf((*MockProductControllerModel)(nil).ReadedProduct), arg0)
}

// UpdatedProduct mocks base method
func (m *MockProductControllerModel) UpdatedProduct(arg0 json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedProduct", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*jrpc2.ErrorObject)
	return ret0, ret1
}

// UpdatedProduct indicates an expected call of UpdatedProduct
func (mr *MockProductControllerModelMockRecorder) UpdatedProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedProduct", reflect.TypeOf((*MockProductControllerModel)(nil).UpdatedProduct), arg0)
}