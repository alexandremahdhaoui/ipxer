// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockcontroller

import (
	controller "github.com/alexandremahdhaoui/ipxer/internal/controller"
	mock "github.com/stretchr/testify/mock"
)

// MockResolveTransformBatchOption is an autogenerated mock type for the ResolveTransformBatchOption type
type MockResolveTransformBatchOption struct {
	mock.Mock
}

type MockResolveTransformBatchOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResolveTransformBatchOption) EXPECT() *MockResolveTransformBatchOption_Expecter {
	return &MockResolveTransformBatchOption_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: options
func (_m *MockResolveTransformBatchOption) Execute(options *controller.ResolveTransformBatchOptions) {
	_m.Called(options)
}

// MockResolveTransformBatchOption_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockResolveTransformBatchOption_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - options *controller.ResolveTransformBatchOptions
func (_e *MockResolveTransformBatchOption_Expecter) Execute(options interface{}) *MockResolveTransformBatchOption_Execute_Call {
	return &MockResolveTransformBatchOption_Execute_Call{Call: _e.mock.On("Execute", options)}
}

func (_c *MockResolveTransformBatchOption_Execute_Call) Run(run func(options *controller.ResolveTransformBatchOptions)) *MockResolveTransformBatchOption_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*controller.ResolveTransformBatchOptions))
	})
	return _c
}

func (_c *MockResolveTransformBatchOption_Execute_Call) Return() *MockResolveTransformBatchOption_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockResolveTransformBatchOption_Execute_Call) RunAndReturn(run func(*controller.ResolveTransformBatchOptions)) *MockResolveTransformBatchOption_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResolveTransformBatchOption creates a new instance of MockResolveTransformBatchOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResolveTransformBatchOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResolveTransformBatchOption {
	mock := &MockResolveTransformBatchOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
