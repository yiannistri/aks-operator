// Code generated by MockGen. DO NOT EDIT.
// Source: ../groups.go
//
// Generated by this command:
//
//	mockgen -destination groups_mock.go -package mock_services -source ../groups.go ResourceGroupsClientInterface
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	armresources "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	gomock "go.uber.org/mock/gomock"
)

// MockResourceGroupsClientInterface is a mock of ResourceGroupsClientInterface interface.
type MockResourceGroupsClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceGroupsClientInterfaceMockRecorder
}

// MockResourceGroupsClientInterfaceMockRecorder is the mock recorder for MockResourceGroupsClientInterface.
type MockResourceGroupsClientInterfaceMockRecorder struct {
	mock *MockResourceGroupsClientInterface
}

// NewMockResourceGroupsClientInterface creates a new mock instance.
func NewMockResourceGroupsClientInterface(ctrl *gomock.Controller) *MockResourceGroupsClientInterface {
	mock := &MockResourceGroupsClientInterface{ctrl: ctrl}
	mock.recorder = &MockResourceGroupsClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceGroupsClientInterface) EXPECT() *MockResourceGroupsClientInterfaceMockRecorder {
	return m.recorder
}

// CheckExistence mocks base method.
func (m *MockResourceGroupsClientInterface) CheckExistence(ctx context.Context, resourceGroupName string) (armresources.ResourceGroupsClientCheckExistenceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExistence", ctx, resourceGroupName)
	ret0, _ := ret[0].(armresources.ResourceGroupsClientCheckExistenceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckExistence indicates an expected call of CheckExistence.
func (mr *MockResourceGroupsClientInterfaceMockRecorder) CheckExistence(ctx, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExistence", reflect.TypeOf((*MockResourceGroupsClientInterface)(nil).CheckExistence), ctx, resourceGroupName)
}

// CreateOrUpdate mocks base method.
func (m *MockResourceGroupsClientInterface) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceGroup armresources.ResourceGroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceGroup)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockResourceGroupsClientInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceGroup any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockResourceGroupsClientInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceGroup)
}
