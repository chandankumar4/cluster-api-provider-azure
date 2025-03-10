/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../managedclusters.go
//
// Generated by this command:
//
//	mockgen -destination managedclusters_mock.go -package mock_managedclusters -source ../managedclusters.go ManagedClusterScope
//
// Package mock_managedclusters is a generated GoMock package.
package mock_managedclusters

import (
	context "context"
	reflect "reflect"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockManagedClusterScope is a mock of ManagedClusterScope interface.
type MockManagedClusterScope struct {
	ctrl     *gomock.Controller
	recorder *MockManagedClusterScopeMockRecorder
}

// MockManagedClusterScopeMockRecorder is the mock recorder for MockManagedClusterScope.
type MockManagedClusterScopeMockRecorder struct {
	mock *MockManagedClusterScope
}

// NewMockManagedClusterScope creates a new mock instance.
func NewMockManagedClusterScope(ctrl *gomock.Controller) *MockManagedClusterScope {
	mock := &MockManagedClusterScope{ctrl: ctrl}
	mock.recorder = &MockManagedClusterScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManagedClusterScope) EXPECT() *MockManagedClusterScopeMockRecorder {
	return m.recorder
}

// AreLocalAccountsDisabled mocks base method.
func (m *MockManagedClusterScope) AreLocalAccountsDisabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreLocalAccountsDisabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AreLocalAccountsDisabled indicates an expected call of AreLocalAccountsDisabled.
func (mr *MockManagedClusterScopeMockRecorder) AreLocalAccountsDisabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreLocalAccountsDisabled", reflect.TypeOf((*MockManagedClusterScope)(nil).AreLocalAccountsDisabled))
}

// BaseURI mocks base method.
func (m *MockManagedClusterScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockManagedClusterScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockManagedClusterScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockManagedClusterScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockManagedClusterScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockManagedClusterScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockManagedClusterScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockManagedClusterScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockManagedClusterScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockManagedClusterScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockManagedClusterScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockManagedClusterScope)(nil).CloudEnvironment))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockManagedClusterScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockManagedClusterScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockManagedClusterScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// GetAdminKubeconfigData mocks base method.
func (m *MockManagedClusterScope) GetAdminKubeconfigData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminKubeconfigData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetAdminKubeconfigData indicates an expected call of GetAdminKubeconfigData.
func (mr *MockManagedClusterScopeMockRecorder) GetAdminKubeconfigData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminKubeconfigData", reflect.TypeOf((*MockManagedClusterScope)(nil).GetAdminKubeconfigData))
}

// GetLongRunningOperationState mocks base method.
func (m *MockManagedClusterScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockManagedClusterScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockManagedClusterScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// GetUserKubeconfigData mocks base method.
func (m *MockManagedClusterScope) GetUserKubeconfigData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserKubeconfigData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetUserKubeconfigData indicates an expected call of GetUserKubeconfigData.
func (mr *MockManagedClusterScopeMockRecorder) GetUserKubeconfigData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserKubeconfigData", reflect.TypeOf((*MockManagedClusterScope)(nil).GetUserKubeconfigData))
}

// HashKey mocks base method.
func (m *MockManagedClusterScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockManagedClusterScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockManagedClusterScope)(nil).HashKey))
}

// IsAADEnabled mocks base method.
func (m *MockManagedClusterScope) IsAADEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAADEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAADEnabled indicates an expected call of IsAADEnabled.
func (mr *MockManagedClusterScopeMockRecorder) IsAADEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAADEnabled", reflect.TypeOf((*MockManagedClusterScope)(nil).IsAADEnabled))
}

// MakeClusterCA mocks base method.
func (m *MockManagedClusterScope) MakeClusterCA() *v1.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeClusterCA")
	ret0, _ := ret[0].(*v1.Secret)
	return ret0
}

// MakeClusterCA indicates an expected call of MakeClusterCA.
func (mr *MockManagedClusterScopeMockRecorder) MakeClusterCA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeClusterCA", reflect.TypeOf((*MockManagedClusterScope)(nil).MakeClusterCA))
}

// MakeEmptyKubeConfigSecret mocks base method.
func (m *MockManagedClusterScope) MakeEmptyKubeConfigSecret() v1.Secret {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeEmptyKubeConfigSecret")
	ret0, _ := ret[0].(v1.Secret)
	return ret0
}

// MakeEmptyKubeConfigSecret indicates an expected call of MakeEmptyKubeConfigSecret.
func (mr *MockManagedClusterScopeMockRecorder) MakeEmptyKubeConfigSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeEmptyKubeConfigSecret", reflect.TypeOf((*MockManagedClusterScope)(nil).MakeEmptyKubeConfigSecret))
}

// ManagedClusterSpec mocks base method.
func (m *MockManagedClusterScope) ManagedClusterSpec() azure.ResourceSpecGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManagedClusterSpec")
	ret0, _ := ret[0].(azure.ResourceSpecGetter)
	return ret0
}

// ManagedClusterSpec indicates an expected call of ManagedClusterSpec.
func (mr *MockManagedClusterScopeMockRecorder) ManagedClusterSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagedClusterSpec", reflect.TypeOf((*MockManagedClusterScope)(nil).ManagedClusterSpec))
}

// SetAdminKubeconfigData mocks base method.
func (m *MockManagedClusterScope) SetAdminKubeconfigData(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAdminKubeconfigData", arg0)
}

// SetAdminKubeconfigData indicates an expected call of SetAdminKubeconfigData.
func (mr *MockManagedClusterScopeMockRecorder) SetAdminKubeconfigData(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAdminKubeconfigData", reflect.TypeOf((*MockManagedClusterScope)(nil).SetAdminKubeconfigData), arg0)
}

// SetControlPlaneEndpoint mocks base method.
func (m *MockManagedClusterScope) SetControlPlaneEndpoint(arg0 v1beta10.APIEndpoint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetControlPlaneEndpoint", arg0)
}

// SetControlPlaneEndpoint indicates an expected call of SetControlPlaneEndpoint.
func (mr *MockManagedClusterScopeMockRecorder) SetControlPlaneEndpoint(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControlPlaneEndpoint", reflect.TypeOf((*MockManagedClusterScope)(nil).SetControlPlaneEndpoint), arg0)
}

// SetKubeletIdentity mocks base method.
func (m *MockManagedClusterScope) SetKubeletIdentity(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKubeletIdentity", arg0)
}

// SetKubeletIdentity indicates an expected call of SetKubeletIdentity.
func (mr *MockManagedClusterScopeMockRecorder) SetKubeletIdentity(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKubeletIdentity", reflect.TypeOf((*MockManagedClusterScope)(nil).SetKubeletIdentity), arg0)
}

// SetLongRunningOperationState mocks base method.
func (m *MockManagedClusterScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockManagedClusterScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockManagedClusterScope)(nil).SetLongRunningOperationState), arg0)
}

// SetOIDCIssuerProfileStatus mocks base method.
func (m *MockManagedClusterScope) SetOIDCIssuerProfileStatus(arg0 *v1beta1.OIDCIssuerProfileStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOIDCIssuerProfileStatus", arg0)
}

// SetOIDCIssuerProfileStatus indicates an expected call of SetOIDCIssuerProfileStatus.
func (mr *MockManagedClusterScopeMockRecorder) SetOIDCIssuerProfileStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOIDCIssuerProfileStatus", reflect.TypeOf((*MockManagedClusterScope)(nil).SetOIDCIssuerProfileStatus), arg0)
}

// SetUserKubeconfigData mocks base method.
func (m *MockManagedClusterScope) SetUserKubeconfigData(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserKubeconfigData", arg0)
}

// SetUserKubeconfigData indicates an expected call of SetUserKubeconfigData.
func (mr *MockManagedClusterScopeMockRecorder) SetUserKubeconfigData(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserKubeconfigData", reflect.TypeOf((*MockManagedClusterScope)(nil).SetUserKubeconfigData), arg0)
}

// StoreClusterInfo mocks base method.
func (m *MockManagedClusterScope) StoreClusterInfo(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreClusterInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreClusterInfo indicates an expected call of StoreClusterInfo.
func (mr *MockManagedClusterScopeMockRecorder) StoreClusterInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreClusterInfo", reflect.TypeOf((*MockManagedClusterScope)(nil).StoreClusterInfo), arg0, arg1)
}

// SubscriptionID mocks base method.
func (m *MockManagedClusterScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockManagedClusterScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockManagedClusterScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockManagedClusterScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockManagedClusterScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockManagedClusterScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockManagedClusterScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockManagedClusterScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockManagedClusterScope)(nil).Token))
}

// UpdateDeleteStatus mocks base method.
func (m *MockManagedClusterScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockManagedClusterScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockManagedClusterScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockManagedClusterScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockManagedClusterScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockManagedClusterScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockManagedClusterScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockManagedClusterScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockManagedClusterScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}
