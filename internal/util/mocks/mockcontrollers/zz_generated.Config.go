// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockcontrollers

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockConfig is an autogenerated mock type for the Config type
type MockConfig struct {
	mock.Mock
}

type MockConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConfig) EXPECT() *MockConfig_Expecter {
	return &MockConfig_Expecter{mock: &_m.Mock}
}

// GetByID provides a mock function with given fields: ctx, profileName, configID
func (_m *MockConfig) GetByID(ctx context.Context, profileName string, configID uuid.UUID) ([]byte, error) {
	ret := _m.Called(ctx, profileName, configID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) ([]byte, error)); ok {
		return rf(ctx, profileName, configID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) []byte); ok {
		r0 = rf(ctx, profileName, configID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, uuid.UUID) error); ok {
		r1 = rf(ctx, profileName, configID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConfig_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockConfig_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - profileName string
//   - configID uuid.UUID
func (_e *MockConfig_Expecter) GetByID(ctx interface{}, profileName interface{}, configID interface{}) *MockConfig_GetByID_Call {
	return &MockConfig_GetByID_Call{Call: _e.mock.On("GetByID", ctx, profileName, configID)}
}

func (_c *MockConfig_GetByID_Call) Run(run func(ctx context.Context, profileName string, configID uuid.UUID)) *MockConfig_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *MockConfig_GetByID_Call) Return(_a0 []byte, _a1 error) *MockConfig_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConfig_GetByID_Call) RunAndReturn(run func(context.Context, string, uuid.UUID) ([]byte, error)) *MockConfig_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConfig creates a new instance of MockConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConfig {
	mock := &MockConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
