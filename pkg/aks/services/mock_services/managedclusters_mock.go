// Code generated by MockGen. DO NOT EDIT.
// Source: ../managedclusters.go
//
// Generated by this command:
//
//	mockgen -destination managedclusters_mock.go -package mock_services -source ../managedclusters.go ManagedClustersClientInterface
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
	services "github.com/rancher/aks-operator/pkg/aks/services"
	gomock "go.uber.org/mock/gomock"
)

// MockPoller is a mock of Poller interface.
type MockPoller[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockPollerMockRecorder[T]
}

// MockPollerMockRecorder is the mock recorder for MockPoller.
type MockPollerMockRecorder[T any] struct {
	mock *MockPoller[T]
}

// NewMockPoller creates a new mock instance.
func NewMockPoller[T any](ctrl *gomock.Controller) *MockPoller[T] {
	mock := &MockPoller[T]{ctrl: ctrl}
	mock.recorder = &MockPollerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPoller[T]) EXPECT() *MockPollerMockRecorder[T] {
	return m.recorder
}

// PollUntilDone mocks base method.
func (m *MockPoller[T]) PollUntilDone(ctx context.Context, options *runtime.PollUntilDoneOptions) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PollUntilDone", ctx, options)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollUntilDone indicates an expected call of PollUntilDone.
func (mr *MockPollerMockRecorder[T]) PollUntilDone(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollUntilDone", reflect.TypeOf((*MockPoller[T])(nil).PollUntilDone), ctx, options)
}

// MockManagedClustersClientInterface is a mock of ManagedClustersClientInterface interface.
type MockManagedClustersClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockManagedClustersClientInterfaceMockRecorder
}

// MockManagedClustersClientInterfaceMockRecorder is the mock recorder for MockManagedClustersClientInterface.
type MockManagedClustersClientInterfaceMockRecorder struct {
	mock *MockManagedClustersClientInterface
}

// NewMockManagedClustersClientInterface creates a new mock instance.
func NewMockManagedClustersClientInterface(ctrl *gomock.Controller) *MockManagedClustersClientInterface {
	mock := &MockManagedClustersClientInterface{ctrl: ctrl}
	mock.recorder = &MockManagedClustersClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagedClustersClientInterface) EXPECT() *MockManagedClustersClientInterfaceMockRecorder {
	return m.recorder
}

// BeginCreateOrUpdate mocks base method.
func (m *MockManagedClustersClientInterface) BeginCreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, parameters armcontainerservice.ManagedCluster, options *armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions) (services.Poller[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginCreateOrUpdate", ctx, resourceGroupName, resourceName, parameters, options)
	ret0, _ := ret[0].(services.Poller[armcontainerservice.ManagedClustersClientCreateOrUpdateResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginCreateOrUpdate indicates an expected call of BeginCreateOrUpdate.
func (mr *MockManagedClustersClientInterfaceMockRecorder) BeginCreateOrUpdate(ctx, resourceGroupName, resourceName, parameters, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginCreateOrUpdate", reflect.TypeOf((*MockManagedClustersClientInterface)(nil).BeginCreateOrUpdate), ctx, resourceGroupName, resourceName, parameters, options)
}

// BeginDelete mocks base method.
func (m *MockManagedClustersClientInterface) BeginDelete(ctx context.Context, resourceGroupName, resourceName string, options *armcontainerservice.ManagedClustersClientBeginDeleteOptions) (services.Poller[armcontainerservice.ManagedClustersClientDeleteResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginDelete", ctx, resourceGroupName, resourceName, options)
	ret0, _ := ret[0].(services.Poller[armcontainerservice.ManagedClustersClientDeleteResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginDelete indicates an expected call of BeginDelete.
func (mr *MockManagedClustersClientInterfaceMockRecorder) BeginDelete(ctx, resourceGroupName, resourceName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginDelete", reflect.TypeOf((*MockManagedClustersClientInterface)(nil).BeginDelete), ctx, resourceGroupName, resourceName, options)
}

// BeginUpdateTags mocks base method.
func (m *MockManagedClustersClientInterface) BeginUpdateTags(ctx context.Context, resourceGroupName, resourceName string, parameters armcontainerservice.TagsObject, options *armcontainerservice.ManagedClustersClientBeginUpdateTagsOptions) (services.Poller[armcontainerservice.ManagedClustersClientUpdateTagsResponse], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginUpdateTags", ctx, resourceGroupName, resourceName, parameters, options)
	ret0, _ := ret[0].(services.Poller[armcontainerservice.ManagedClustersClientUpdateTagsResponse])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginUpdateTags indicates an expected call of BeginUpdateTags.
func (mr *MockManagedClustersClientInterfaceMockRecorder) BeginUpdateTags(ctx, resourceGroupName, resourceName, parameters, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginUpdateTags", reflect.TypeOf((*MockManagedClustersClientInterface)(nil).BeginUpdateTags), ctx, resourceGroupName, resourceName, parameters, options)
}

// Get mocks base method.
func (m *MockManagedClustersClientInterface) Get(ctx context.Context, resourceGroupName, resourceName string, options *armcontainerservice.ManagedClustersClientGetOptions) (armcontainerservice.ManagedClustersClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, options)
	ret0, _ := ret[0].(armcontainerservice.ManagedClustersClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockManagedClustersClientInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManagedClustersClientInterface)(nil).Get), ctx, resourceGroupName, resourceName, options)
}

// GetAccessProfile mocks base method.
func (m *MockManagedClustersClientInterface) GetAccessProfile(ctx context.Context, resourceGroupName, resourceName, roleName string, options *armcontainerservice.ManagedClustersClientGetAccessProfileOptions) (armcontainerservice.ManagedClustersClientGetAccessProfileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessProfile", ctx, resourceGroupName, resourceName, roleName, options)
	ret0, _ := ret[0].(armcontainerservice.ManagedClustersClientGetAccessProfileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessProfile indicates an expected call of GetAccessProfile.
func (mr *MockManagedClustersClientInterfaceMockRecorder) GetAccessProfile(ctx, resourceGroupName, resourceName, roleName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessProfile", reflect.TypeOf((*MockManagedClustersClientInterface)(nil).GetAccessProfile), ctx, resourceGroupName, resourceName, roleName, options)
}
