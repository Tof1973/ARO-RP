// Code generated by MockGen. DO NOT EDIT.
// Source: ./federated_identity_credentials.go
//
// Generated by this command:
//
//	mockgen -source ./federated_identity_credentials.go -destination=../../../mocks/azureclient/azuresdk/armmsi/federated_identity_credentials.go github.com/Azure/ARO-RP/pkg/util/azureclient/azuresdk/armmsi FederatedIdentityCredentialsClient
//

// Package mock_armmsi is a generated GoMock package.
package mock_armmsi

import (
	context "context"
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armmsi "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
	gomock "go.uber.org/mock/gomock"
)

// MockFederatedIdentityCredentialsClient is a mock of FederatedIdentityCredentialsClient interface.
type MockFederatedIdentityCredentialsClient struct {
	ctrl     *gomock.Controller
	recorder *MockFederatedIdentityCredentialsClientMockRecorder
}

// MockFederatedIdentityCredentialsClientMockRecorder is the mock recorder for MockFederatedIdentityCredentialsClient.
type MockFederatedIdentityCredentialsClientMockRecorder struct {
	mock *MockFederatedIdentityCredentialsClient
}

// NewMockFederatedIdentityCredentialsClient creates a new mock instance.
func NewMockFederatedIdentityCredentialsClient(ctrl *gomock.Controller) *MockFederatedIdentityCredentialsClient {
	mock := &MockFederatedIdentityCredentialsClient{ctrl: ctrl}
	mock.recorder = &MockFederatedIdentityCredentialsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFederatedIdentityCredentialsClient) EXPECT() *MockFederatedIdentityCredentialsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockFederatedIdentityCredentialsClient) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName, federatedIdentityCredentialResourceName string, parameters armmsi.FederatedIdentityCredential, options *armmsi.FederatedIdentityCredentialsClientCreateOrUpdateOptions) (armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options)
	ret0, _ := ret[0].(armmsi.FederatedIdentityCredentialsClientCreateOrUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options)
}

// Delete mocks base method.
func (m *MockFederatedIdentityCredentialsClient) Delete(ctx context.Context, resourceGroupName, resourceName, federatedIdentityCredentialResourceName string, options *armmsi.FederatedIdentityCredentialsClientDeleteOptions) (armmsi.FederatedIdentityCredentialsClientDeleteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
	ret0, _ := ret[0].(armmsi.FederatedIdentityCredentialsClientDeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) Delete(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).Delete), ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
}

// Get mocks base method.
func (m *MockFederatedIdentityCredentialsClient) Get(ctx context.Context, resourceGroupName, resourceName, federatedIdentityCredentialResourceName string, options *armmsi.FederatedIdentityCredentialsClientGetOptions) (armmsi.FederatedIdentityCredentialsClientGetResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
	ret0, _ := ret[0].(armmsi.FederatedIdentityCredentialsClientGetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) Get(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).Get), ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
}

// NewListPager mocks base method.
func (m *MockFederatedIdentityCredentialsClient) NewListPager(resourceGroupName, resourceName string, options *armmsi.FederatedIdentityCredentialsClientListOptions) *runtime.Pager[armmsi.FederatedIdentityCredentialsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", resourceGroupName, resourceName, options)
	ret0, _ := ret[0].(*runtime.Pager[armmsi.FederatedIdentityCredentialsClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockFederatedIdentityCredentialsClientMockRecorder) NewListPager(resourceGroupName, resourceName, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockFederatedIdentityCredentialsClient)(nil).NewListPager), resourceGroupName, resourceName, options)
}