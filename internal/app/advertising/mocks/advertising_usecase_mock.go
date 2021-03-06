// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package advertising_mock is a generated GoMock package.
package advertising_mock

import (
	advertisingModel "github.com/Kostikans/AvitoTechadvertising/internal/app/advertising/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUsecase is a mock of Usecase interface
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// AddAdvertising mocks base method
func (m *MockUsecase) AddAdvertising(advertising advertisingModel.AdvertisingAdd) (advertisingModel.AdvertisingID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAdvertising", advertising)
	ret0, _ := ret[0].(advertisingModel.AdvertisingID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAdvertising indicates an expected call of AddAdvertising
func (mr *MockUsecaseMockRecorder) AddAdvertising(advertising interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdvertising", reflect.TypeOf((*MockUsecase)(nil).AddAdvertising), advertising)
}

// GetAdvertising mocks base method
func (m *MockUsecase) GetAdvertising(advertisingID int, fields string) (advertisingModel.Advertising, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdvertising", advertisingID, fields)
	ret0, _ := ret[0].(advertisingModel.Advertising)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdvertising indicates an expected call of GetAdvertising
func (mr *MockUsecaseMockRecorder) GetAdvertising(advertisingID, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdvertising", reflect.TypeOf((*MockUsecase)(nil).GetAdvertising), advertisingID, fields)
}

// ListAdvertising mocks base method
func (m *MockUsecase) ListAdvertising(field, desc string, page int) (advertisingModel.AdvertisingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdvertising", field, desc, page)
	ret0, _ := ret[0].(advertisingModel.AdvertisingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAdvertising indicates an expected call of ListAdvertising
func (mr *MockUsecaseMockRecorder) ListAdvertising(field, desc, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdvertising", reflect.TypeOf((*MockUsecase)(nil).ListAdvertising), field, desc, page)
}

// CheckAdvertisingExist mocks base method
func (m *MockUsecase) CheckAdvertisingExist(AdvertisingID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAdvertisingExist", AdvertisingID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAdvertisingExist indicates an expected call of CheckAdvertisingExist
func (mr *MockUsecaseMockRecorder) CheckAdvertisingExist(AdvertisingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAdvertisingExist", reflect.TypeOf((*MockUsecase)(nil).CheckAdvertisingExist), AdvertisingID)
}
