// Code generated by MockGen. DO NOT EDIT.
// Source: ../workplaces.go
//
// Generated by this command:
//
//	mockgen -destination workplaces_mock.go -package mock_services -source ../workplaces.go WorkplacesClientInterface
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	operationalinsights "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-08-01/operationalinsights"
	gomock "go.uber.org/mock/gomock"
)

// MockWorkplacesClientInterface is a mock of WorkplacesClientInterface interface.
type MockWorkplacesClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockWorkplacesClientInterfaceMockRecorder
}

// MockWorkplacesClientInterfaceMockRecorder is the mock recorder for MockWorkplacesClientInterface.
type MockWorkplacesClientInterfaceMockRecorder struct {
	mock *MockWorkplacesClientInterface
}

// NewMockWorkplacesClientInterface creates a new mock instance.
func NewMockWorkplacesClientInterface(ctrl *gomock.Controller) *MockWorkplacesClientInterface {
	mock := &MockWorkplacesClientInterface{ctrl: ctrl}
	mock.recorder = &MockWorkplacesClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkplacesClientInterface) EXPECT() *MockWorkplacesClientInterfaceMockRecorder {
	return m.recorder
}

// AsyncCreateUpdateResult mocks base method.
func (m *MockWorkplacesClientInterface) AsyncCreateUpdateResult(asyncRet operationalinsights.WorkspacesCreateOrUpdateFuture) (operationalinsights.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncCreateUpdateResult", asyncRet)
	ret0, _ := ret[0].(operationalinsights.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AsyncCreateUpdateResult indicates an expected call of AsyncCreateUpdateResult.
func (mr *MockWorkplacesClientInterfaceMockRecorder) AsyncCreateUpdateResult(asyncRet any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncCreateUpdateResult", reflect.TypeOf((*MockWorkplacesClientInterface)(nil).AsyncCreateUpdateResult), asyncRet)
}

// CreateOrUpdate mocks base method.
func (m *MockWorkplacesClientInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, workspaceName string, parameters operationalinsights.Workspace) (operationalinsights.WorkspacesCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, workspaceName, parameters)
	ret0, _ := ret[0].(operationalinsights.WorkspacesCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockWorkplacesClientInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, workspaceName, parameters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockWorkplacesClientInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, workspaceName, parameters)
}

// Get mocks base method.
func (m *MockWorkplacesClientInterface) Get(ctx context.Context, resourceGroupName, workspaceName string) (operationalinsights.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, workspaceName)
	ret0, _ := ret[0].(operationalinsights.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockWorkplacesClientInterfaceMockRecorder) Get(ctx, resourceGroupName, workspaceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWorkplacesClientInterface)(nil).Get), ctx, resourceGroupName, workspaceName)
}
