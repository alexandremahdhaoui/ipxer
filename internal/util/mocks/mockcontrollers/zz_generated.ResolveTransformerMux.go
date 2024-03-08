// Code generated by mockery v2.42.0. DO NOT EDIT.

package mockcontrollers

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/alexandremahdhaoui/ipxer/internal/types"
)

// MockResolveTransformerMux is an autogenerated mock type for the ResolveTransformerMux type
type MockResolveTransformerMux struct {
	mock.Mock
}

type MockResolveTransformerMux_Expecter struct {
	mock *mock.Mock
}

func (_m *MockResolveTransformerMux) EXPECT() *MockResolveTransformerMux_Expecter {
	return &MockResolveTransformerMux_Expecter{mock: &_m.Mock}
}

// ResolveAndTransformBatch provides a mock function with given fields: ctx, batch
func (_m *MockResolveTransformerMux) ResolveAndTransformBatch(ctx context.Context, batch []types.Content) (map[string][]byte, error) {
	ret := _m.Called(ctx, batch)

	if len(ret) == 0 {
		panic("no return value specified for ResolveAndTransformBatch")
	}

	var r0 map[string][]byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.Content) (map[string][]byte, error)); ok {
		return rf(ctx, batch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.Content) map[string][]byte); ok {
		r0 = rf(ctx, batch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.Content) error); ok {
		r1 = rf(ctx, batch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockResolveTransformerMux_ResolveAndTransformBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ResolveAndTransformBatch'
type MockResolveTransformerMux_ResolveAndTransformBatch_Call struct {
	*mock.Call
}

// ResolveAndTransformBatch is a helper method to define mock.On call
//   - ctx context.Context
//   - batch []types.Content
func (_e *MockResolveTransformerMux_Expecter) ResolveAndTransformBatch(ctx interface{}, batch interface{}) *MockResolveTransformerMux_ResolveAndTransformBatch_Call {
	return &MockResolveTransformerMux_ResolveAndTransformBatch_Call{Call: _e.mock.On("ResolveAndTransformBatch", ctx, batch)}
}

func (_c *MockResolveTransformerMux_ResolveAndTransformBatch_Call) Run(run func(ctx context.Context, batch []types.Content)) *MockResolveTransformerMux_ResolveAndTransformBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]types.Content))
	})
	return _c
}

func (_c *MockResolveTransformerMux_ResolveAndTransformBatch_Call) Return(_a0 map[string][]byte, _a1 error) *MockResolveTransformerMux_ResolveAndTransformBatch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockResolveTransformerMux_ResolveAndTransformBatch_Call) RunAndReturn(run func(context.Context, []types.Content) (map[string][]byte, error)) *MockResolveTransformerMux_ResolveAndTransformBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockResolveTransformerMux creates a new instance of MockResolveTransformerMux. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResolveTransformerMux(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResolveTransformerMux {
	mock := &MockResolveTransformerMux{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
