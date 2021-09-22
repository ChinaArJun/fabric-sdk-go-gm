// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ChinaArJun/fabric-sdk-go-gm/pkg/common/providers/core (interfaces: CryptoSuiteConfig,ConfigBackend,Providers)

// Package mockcore is a generated GoMock package.
package mockcore

import (
	reflect "reflect"

	core "github.com/ChinaArJun/fabric-sdk-go-gm/pkg/common/providers/core"
	gomock "github.com/golang/mock/gomock"
)

// MockCryptoSuiteConfig is a mock of CryptoSuiteConfig interface
type MockCryptoSuiteConfig struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoSuiteConfigMockRecorder
}

// MockCryptoSuiteConfigMockRecorder is the mock recorder for MockCryptoSuiteConfig
type MockCryptoSuiteConfigMockRecorder struct {
	mock *MockCryptoSuiteConfig
}

// NewMockCryptoSuiteConfig creates a new mock instance
func NewMockCryptoSuiteConfig(ctrl *gomock.Controller) *MockCryptoSuiteConfig {
	mock := &MockCryptoSuiteConfig{ctrl: ctrl}
	mock.recorder = &MockCryptoSuiteConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCryptoSuiteConfig) EXPECT() *MockCryptoSuiteConfigMockRecorder {
	return m.recorder
}

// IsSecurityEnabled mocks base method
func (m *MockCryptoSuiteConfig) IsSecurityEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSecurityEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecurityEnabled indicates an expected call of IsSecurityEnabled
func (mr *MockCryptoSuiteConfigMockRecorder) IsSecurityEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSecurityEnabled", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).IsSecurityEnabled))
}

// KeyStorePath mocks base method
func (m *MockCryptoSuiteConfig) KeyStorePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// KeyStorePath indicates an expected call of KeyStorePath
func (mr *MockCryptoSuiteConfigMockRecorder) KeyStorePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyStorePath", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).KeyStorePath))
}

// SecurityAlgorithm mocks base method
func (m *MockCryptoSuiteConfig) SecurityAlgorithm() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityAlgorithm indicates an expected call of SecurityAlgorithm
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityAlgorithm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityAlgorithm", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityAlgorithm))
}

// SecurityLevel mocks base method
func (m *MockCryptoSuiteConfig) SecurityLevel() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// SecurityLevel indicates an expected call of SecurityLevel
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityLevel", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityLevel))
}

// SecurityProvider mocks base method
func (m *MockCryptoSuiteConfig) SecurityProvider() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityProvider")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProvider indicates an expected call of SecurityProvider
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProvider", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProvider))
}

// SecurityProviderLabel mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderLabel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityProviderLabel")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderLabel indicates an expected call of SecurityProviderLabel
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderLabel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderLabel", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderLabel))
}

// SecurityProviderLibPath mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderLibPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityProviderLibPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderLibPath indicates an expected call of SecurityProviderLibPath
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderLibPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderLibPath", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderLibPath))
}

// SecurityProviderPin mocks base method
func (m *MockCryptoSuiteConfig) SecurityProviderPin() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityProviderPin")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderPin indicates an expected call of SecurityProviderPin
func (mr *MockCryptoSuiteConfigMockRecorder) SecurityProviderPin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProviderPin", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SecurityProviderPin))
}

// SoftVerify mocks base method
func (m *MockCryptoSuiteConfig) SoftVerify() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftVerify")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SoftVerify indicates an expected call of SoftVerify
func (mr *MockCryptoSuiteConfigMockRecorder) SoftVerify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftVerify", reflect.TypeOf((*MockCryptoSuiteConfig)(nil).SoftVerify))
}

// MockConfigBackend is a mock of ConfigBackend interface
type MockConfigBackend struct {
	ctrl     *gomock.Controller
	recorder *MockConfigBackendMockRecorder
}

// MockConfigBackendMockRecorder is the mock recorder for MockConfigBackend
type MockConfigBackendMockRecorder struct {
	mock *MockConfigBackend
}

// NewMockConfigBackend creates a new mock instance
func NewMockConfigBackend(ctrl *gomock.Controller) *MockConfigBackend {
	mock := &MockConfigBackend{ctrl: ctrl}
	mock.recorder = &MockConfigBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigBackend) EXPECT() *MockConfigBackendMockRecorder {
	return m.recorder
}

// Lookup mocks base method
func (m *MockConfigBackend) Lookup(arg0 string) (interface{}, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lookup", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Lookup indicates an expected call of Lookup
func (mr *MockConfigBackendMockRecorder) Lookup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lookup", reflect.TypeOf((*MockConfigBackend)(nil).Lookup), arg0)
}

// MockProviders is a mock of Providers interface
type MockProviders struct {
	ctrl     *gomock.Controller
	recorder *MockProvidersMockRecorder
}

// MockProvidersMockRecorder is the mock recorder for MockProviders
type MockProvidersMockRecorder struct {
	mock *MockProviders
}

// NewMockProviders creates a new mock instance
func NewMockProviders(ctrl *gomock.Controller) *MockProviders {
	mock := &MockProviders{ctrl: ctrl}
	mock.recorder = &MockProvidersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProviders) EXPECT() *MockProvidersMockRecorder {
	return m.recorder
}

// CryptoSuite mocks base method
func (m *MockProviders) CryptoSuite() core.CryptoSuite {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CryptoSuite")
	ret0, _ := ret[0].(core.CryptoSuite)
	return ret0
}

// CryptoSuite indicates an expected call of CryptoSuite
func (mr *MockProvidersMockRecorder) CryptoSuite() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptoSuite", reflect.TypeOf((*MockProviders)(nil).CryptoSuite))
}

// SigningManager mocks base method
func (m *MockProviders) SigningManager() core.SigningManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SigningManager")
	ret0, _ := ret[0].(core.SigningManager)
	return ret0
}

// SigningManager indicates an expected call of SigningManager
func (mr *MockProvidersMockRecorder) SigningManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SigningManager", reflect.TypeOf((*MockProviders)(nil).SigningManager))
}
