// Code generated by mockery v2.21.1. DO NOT EDIT.

package mocks

import (
	context "context"

	bbolt "go.etcd.io/bbolt"

	keys "github.com/kolide/launcher/ee/agent/flags/keys"

	mock "github.com/stretchr/testify/mock"

	slog "log/slog"

	storage "github.com/kolide/launcher/ee/agent/storage"

	time "time"

	types "github.com/kolide/launcher/ee/agent/types"
)

// Knapsack is an autogenerated mock type for the Knapsack type
type Knapsack struct {
	mock.Mock
}

// AddSlogHandler provides a mock function with given fields: handler
func (_m *Knapsack) AddSlogHandler(handler ...slog.Handler) {
	_va := make([]interface{}, len(handler))
	for _i := range handler {
		_va[_i] = handler[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// AgentFlagsStore provides a mock function with given fields:
func (_m *Knapsack) AgentFlagsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// Autoupdate provides a mock function with given fields:
func (_m *Knapsack) Autoupdate() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AutoupdateErrorsStore provides a mock function with given fields:
func (_m *Knapsack) AutoupdateErrorsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// AutoupdateInitialDelay provides a mock function with given fields:
func (_m *Knapsack) AutoupdateInitialDelay() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// AutoupdateInterval provides a mock function with given fields:
func (_m *Knapsack) AutoupdateInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// BboltDB provides a mock function with given fields:
func (_m *Knapsack) BboltDB() *bbolt.DB {
	ret := _m.Called()

	var r0 *bbolt.DB
	if rf, ok := ret.Get(0).(func() *bbolt.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bbolt.DB)
		}
	}

	return r0
}

// CertPins provides a mock function with given fields:
func (_m *Knapsack) CertPins() [][]byte {
	ret := _m.Called()

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func() [][]byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	return r0
}

// ConfigStore provides a mock function with given fields:
func (_m *Knapsack) ConfigStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// ControlRequestInterval provides a mock function with given fields:
func (_m *Knapsack) ControlRequestInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// ControlServerURL provides a mock function with given fields:
func (_m *Knapsack) ControlServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ControlStore provides a mock function with given fields:
func (_m *Knapsack) ControlStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// Debug provides a mock function with given fields:
func (_m *Knapsack) Debug() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DebugLogFile provides a mock function with given fields:
func (_m *Knapsack) DebugLogFile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// DebugServerData provides a mock function with given fields:
func (_m *Knapsack) DebugServerData() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DesktopEnabled provides a mock function with given fields:
func (_m *Knapsack) DesktopEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DesktopMenuRefreshInterval provides a mock function with given fields:
func (_m *Knapsack) DesktopMenuRefreshInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// DesktopUpdateInterval provides a mock function with given fields:
func (_m *Knapsack) DesktopUpdateInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// DisableControlTLS provides a mock function with given fields:
func (_m *Knapsack) DisableControlTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// DisableTraceIngestTLS provides a mock function with given fields:
func (_m *Knapsack) DisableTraceIngestTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EnableInitialRunner provides a mock function with given fields:
func (_m *Knapsack) EnableInitialRunner() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EnrollSecret provides a mock function with given fields:
func (_m *Knapsack) EnrollSecret() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EnrollSecretPath provides a mock function with given fields:
func (_m *Knapsack) EnrollSecretPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ExportTraces provides a mock function with given fields:
func (_m *Knapsack) ExportTraces() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ForceControlSubsystems provides a mock function with given fields:
func (_m *Knapsack) ForceControlSubsystems() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IAmBreakingEELicense provides a mock function with given fields:
func (_m *Knapsack) IAmBreakingEELicense() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InModernStandby provides a mock function with given fields:
func (_m *Knapsack) InModernStandby() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InitialResultsStore provides a mock function with given fields:
func (_m *Knapsack) InitialResultsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// InsecureControlTLS provides a mock function with given fields:
func (_m *Knapsack) InsecureControlTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InsecureTLS provides a mock function with given fields:
func (_m *Knapsack) InsecureTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InsecureTransportTLS provides a mock function with given fields:
func (_m *Knapsack) InsecureTransportTLS() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// KolideHosted provides a mock function with given fields:
func (_m *Knapsack) KolideHosted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// KolideServerURL provides a mock function with given fields:
func (_m *Knapsack) KolideServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LatestOsquerydPath provides a mock function with given fields: ctx
func (_m *Knapsack) LatestOsquerydPath(ctx context.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LocalDevelopmentPath provides a mock function with given fields:
func (_m *Knapsack) LocalDevelopmentPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LogIngestServerURL provides a mock function with given fields:
func (_m *Knapsack) LogIngestServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LogMaxBytesPerBatch provides a mock function with given fields:
func (_m *Knapsack) LogMaxBytesPerBatch() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// LogShippingLevel provides a mock function with given fields:
func (_m *Knapsack) LogShippingLevel() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// LoggingInterval provides a mock function with given fields:
func (_m *Knapsack) LoggingInterval() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// MirrorServerURL provides a mock function with given fields:
func (_m *Knapsack) MirrorServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryFlags provides a mock function with given fields:
func (_m *Knapsack) OsqueryFlags() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// OsqueryHealthcheckStartupDelay provides a mock function with given fields:
func (_m *Knapsack) OsqueryHealthcheckStartupDelay() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// OsqueryHistoryInstanceStore provides a mock function with given fields:
func (_m *Knapsack) OsqueryHistoryInstanceStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// OsqueryTlsConfigEndpoint provides a mock function with given fields:
func (_m *Knapsack) OsqueryTlsConfigEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryTlsDistributedReadEndpoint provides a mock function with given fields:
func (_m *Knapsack) OsqueryTlsDistributedReadEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryTlsDistributedWriteEndpoint provides a mock function with given fields:
func (_m *Knapsack) OsqueryTlsDistributedWriteEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryTlsEnrollEndpoint provides a mock function with given fields:
func (_m *Knapsack) OsqueryTlsEnrollEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryTlsLoggerEndpoint provides a mock function with given fields:
func (_m *Knapsack) OsqueryTlsLoggerEndpoint() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OsqueryVerbose provides a mock function with given fields:
func (_m *Knapsack) OsqueryVerbose() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// OsquerydPath provides a mock function with given fields:
func (_m *Knapsack) OsquerydPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PersistentHostDataStore provides a mock function with given fields:
func (_m *Knapsack) PersistentHostDataStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// ReadEnrollSecret provides a mock function with given fields:
func (_m *Knapsack) ReadEnrollSecret() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterChangeObserver provides a mock function with given fields: observer, flagKeys
func (_m *Knapsack) RegisterChangeObserver(observer types.FlagsChangeObserver, flagKeys ...keys.FlagKey) {
	_va := make([]interface{}, len(flagKeys))
	for _i := range flagKeys {
		_va[_i] = flagKeys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, observer)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// ResultLogsStore provides a mock function with given fields:
func (_m *Knapsack) ResultLogsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// RootDirectory provides a mock function with given fields:
func (_m *Knapsack) RootDirectory() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RootPEM provides a mock function with given fields:
func (_m *Knapsack) RootPEM() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SentNotificationsStore provides a mock function with given fields:
func (_m *Knapsack) SentNotificationsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// ServerProvidedDataStore provides a mock function with given fields:
func (_m *Knapsack) ServerProvidedDataStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// SetAutoupdate provides a mock function with given fields: enabled
func (_m *Knapsack) SetAutoupdate(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAutoupdateInitialDelay provides a mock function with given fields: delay
func (_m *Knapsack) SetAutoupdateInitialDelay(delay time.Duration) error {
	ret := _m.Called(delay)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(delay)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetAutoupdateInterval provides a mock function with given fields: interval
func (_m *Knapsack) SetAutoupdateInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetControlRequestInterval provides a mock function with given fields: interval
func (_m *Knapsack) SetControlRequestInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetControlRequestIntervalOverride provides a mock function with given fields: value, duration
func (_m *Knapsack) SetControlRequestIntervalOverride(value time.Duration, duration time.Duration) {
	_m.Called(value, duration)
}

// SetControlServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetControlServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDebug provides a mock function with given fields: debug
func (_m *Knapsack) SetDebug(debug bool) error {
	ret := _m.Called(debug)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(debug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDebugServerData provides a mock function with given fields: debug
func (_m *Knapsack) SetDebugServerData(debug bool) error {
	ret := _m.Called(debug)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(debug)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDesktopEnabled provides a mock function with given fields: enabled
func (_m *Knapsack) SetDesktopEnabled(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDesktopMenuRefreshInterval provides a mock function with given fields: interval
func (_m *Knapsack) SetDesktopMenuRefreshInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDesktopUpdateInterval provides a mock function with given fields: interval
func (_m *Knapsack) SetDesktopUpdateInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDisableControlTLS provides a mock function with given fields: disabled
func (_m *Knapsack) SetDisableControlTLS(disabled bool) error {
	ret := _m.Called(disabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(disabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDisableTraceIngestTLS provides a mock function with given fields: enabled
func (_m *Knapsack) SetDisableTraceIngestTLS(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetExportTraces provides a mock function with given fields: enabled
func (_m *Knapsack) SetExportTraces(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetExportTracesOverride provides a mock function with given fields: value, duration
func (_m *Knapsack) SetExportTracesOverride(value bool, duration time.Duration) {
	_m.Called(value, duration)
}

// SetForceControlSubsystems provides a mock function with given fields: force
func (_m *Knapsack) SetForceControlSubsystems(force bool) error {
	ret := _m.Called(force)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetInModernStandby provides a mock function with given fields: enabled
func (_m *Knapsack) SetInModernStandby(enabled bool) error {
	ret := _m.Called(enabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetInsecureControlTLS provides a mock function with given fields: disabled
func (_m *Knapsack) SetInsecureControlTLS(disabled bool) error {
	ret := _m.Called(disabled)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(disabled)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetInsecureTLS provides a mock function with given fields: insecure
func (_m *Knapsack) SetInsecureTLS(insecure bool) error {
	ret := _m.Called(insecure)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(insecure)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetInsecureTransportTLS provides a mock function with given fields: insecure
func (_m *Knapsack) SetInsecureTransportTLS(insecure bool) error {
	ret := _m.Called(insecure)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(insecure)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetKolideServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetKolideServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogIngestServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetLogIngestServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogShippingLevel provides a mock function with given fields: level
func (_m *Knapsack) SetLogShippingLevel(level string) error {
	ret := _m.Called(level)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(level)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLogShippingLevelOverride provides a mock function with given fields: value, duration
func (_m *Knapsack) SetLogShippingLevelOverride(value string, duration time.Duration) {
	_m.Called(value, duration)
}

// SetLoggingInterval provides a mock function with given fields: interval
func (_m *Knapsack) SetLoggingInterval(interval time.Duration) error {
	ret := _m.Called(interval)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(interval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMirrorServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetMirrorServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOsqueryHealthcheckStartupDelay provides a mock function with given fields: delay
func (_m *Knapsack) SetOsqueryHealthcheckStartupDelay(delay time.Duration) error {
	ret := _m.Called(delay)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(delay)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetOsqueryVerbose provides a mock function with given fields: verbose
func (_m *Knapsack) SetOsqueryVerbose(verbose bool) error {
	ret := _m.Called(verbose)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(verbose)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTraceBatchTimeout provides a mock function with given fields: duration
func (_m *Knapsack) SetTraceBatchTimeout(duration time.Duration) error {
	ret := _m.Called(duration)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration) error); ok {
		r0 = rf(duration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTraceIngestServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetTraceIngestServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTraceSamplingRate provides a mock function with given fields: rate
func (_m *Knapsack) SetTraceSamplingRate(rate float64) error {
	ret := _m.Called(rate)

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(rate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetTraceSamplingRateOverride provides a mock function with given fields: value, duration
func (_m *Knapsack) SetTraceSamplingRateOverride(value float64, duration time.Duration) {
	_m.Called(value, duration)
}

// SetTufServerURL provides a mock function with given fields: url
func (_m *Knapsack) SetTufServerURL(url string) error {
	ret := _m.Called(url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUpdateChannel provides a mock function with given fields: channel
func (_m *Knapsack) SetUpdateChannel(channel string) error {
	ret := _m.Called(channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUpdateDirectory provides a mock function with given fields: directory
func (_m *Knapsack) SetUpdateDirectory(directory string) error {
	ret := _m.Called(directory)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(directory)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWatchdogDelaySec provides a mock function with given fields: sec
func (_m *Knapsack) SetWatchdogDelaySec(sec int) error {
	ret := _m.Called(sec)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(sec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWatchdogEnabled provides a mock function with given fields: enable
func (_m *Knapsack) SetWatchdogEnabled(enable bool) error {
	ret := _m.Called(enable)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(enable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWatchdogMemoryLimitMB provides a mock function with given fields: limit
func (_m *Knapsack) SetWatchdogMemoryLimitMB(limit int) error {
	ret := _m.Called(limit)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWatchdogUtilizationLimitPercent provides a mock function with given fields: limit
func (_m *Knapsack) SetWatchdogUtilizationLimitPercent(limit int) error {
	ret := _m.Called(limit)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Slogger provides a mock function with given fields:
func (_m *Knapsack) Slogger() *slog.Logger {
	ret := _m.Called()

	var r0 *slog.Logger
	if rf, ok := ret.Get(0).(func() *slog.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slog.Logger)
		}
	}

	return r0
}

// StatusLogsStore provides a mock function with given fields:
func (_m *Knapsack) StatusLogsStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// Stores provides a mock function with given fields:
func (_m *Knapsack) Stores() map[storage.Store]types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 map[storage.Store]types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() map[storage.Store]types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[storage.Store]types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// SystemSlogger provides a mock function with given fields:
func (_m *Knapsack) SystemSlogger() *slog.Logger {
	ret := _m.Called()

	var r0 *slog.Logger
	if rf, ok := ret.Get(0).(func() *slog.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*slog.Logger)
		}
	}

	return r0
}

// TokenStore provides a mock function with given fields:
func (_m *Knapsack) TokenStore() types.GetterSetterDeleterIteratorUpdater {
	ret := _m.Called()

	var r0 types.GetterSetterDeleterIteratorUpdater
	if rf, ok := ret.Get(0).(func() types.GetterSetterDeleterIteratorUpdater); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.GetterSetterDeleterIteratorUpdater)
		}
	}

	return r0
}

// TraceBatchTimeout provides a mock function with given fields:
func (_m *Knapsack) TraceBatchTimeout() time.Duration {
	ret := _m.Called()

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// TraceIngestServerURL provides a mock function with given fields:
func (_m *Knapsack) TraceIngestServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TraceSamplingRate provides a mock function with given fields:
func (_m *Knapsack) TraceSamplingRate() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Transport provides a mock function with given fields:
func (_m *Knapsack) Transport() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TufServerURL provides a mock function with given fields:
func (_m *Knapsack) TufServerURL() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateChannel provides a mock function with given fields:
func (_m *Knapsack) UpdateChannel() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateDirectory provides a mock function with given fields:
func (_m *Knapsack) UpdateDirectory() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WatchdogDelaySec provides a mock function with given fields:
func (_m *Knapsack) WatchdogDelaySec() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// WatchdogEnabled provides a mock function with given fields:
func (_m *Knapsack) WatchdogEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// WatchdogMemoryLimitMB provides a mock function with given fields:
func (_m *Knapsack) WatchdogMemoryLimitMB() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// WatchdogUtilizationLimitPercent provides a mock function with given fields:
func (_m *Knapsack) WatchdogUtilizationLimitPercent() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewKnapsack interface {
	mock.TestingT
	Cleanup(func())
}

// NewKnapsack creates a new instance of Knapsack. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKnapsack(t mockConstructorTestingTNewKnapsack) *Knapsack {
	mock := &Knapsack{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
