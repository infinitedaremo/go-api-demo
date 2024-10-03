// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/infinitedaremo/go-api-demo/internal/app"

	mock "github.com/stretchr/testify/mock"
)

// PersonService is an autogenerated mock type for the PersonService type
type PersonService struct {
	mock.Mock
}

type PersonService_Expecter struct {
	mock *mock.Mock
}

func (_m *PersonService) EXPECT() *PersonService_Expecter {
	return &PersonService_Expecter{mock: &_m.Mock}
}

// GetPerson provides a mock function with given fields: _a0, _a1
func (_m *PersonService) GetPerson(_a0 context.Context, _a1 int64) (app.Person, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetPerson")
	}

	var r0 app.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (app.Person, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) app.Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(app.Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonService_GetPerson_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPerson'
type PersonService_GetPerson_Call struct {
	*mock.Call
}

// GetPerson is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int64
func (_e *PersonService_Expecter) GetPerson(_a0 interface{}, _a1 interface{}) *PersonService_GetPerson_Call {
	return &PersonService_GetPerson_Call{Call: _e.mock.On("GetPerson", _a0, _a1)}
}

func (_c *PersonService_GetPerson_Call) Run(run func(_a0 context.Context, _a1 int64)) *PersonService_GetPerson_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *PersonService_GetPerson_Call) Return(_a0 app.Person, _a1 error) *PersonService_GetPerson_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonService_GetPerson_Call) RunAndReturn(run func(context.Context, int64) (app.Person, error)) *PersonService_GetPerson_Call {
	_c.Call.Return(run)
	return _c
}

// GetPortfolio provides a mock function with given fields: _a0, _a1
func (_m *PersonService) GetPortfolio(_a0 context.Context, _a1 int64) (*app.Portfolio, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetPortfolio")
	}

	var r0 *app.Portfolio
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*app.Portfolio, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *app.Portfolio); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app.Portfolio)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonService_GetPortfolio_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPortfolio'
type PersonService_GetPortfolio_Call struct {
	*mock.Call
}

// GetPortfolio is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int64
func (_e *PersonService_Expecter) GetPortfolio(_a0 interface{}, _a1 interface{}) *PersonService_GetPortfolio_Call {
	return &PersonService_GetPortfolio_Call{Call: _e.mock.On("GetPortfolio", _a0, _a1)}
}

func (_c *PersonService_GetPortfolio_Call) Run(run func(_a0 context.Context, _a1 int64)) *PersonService_GetPortfolio_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *PersonService_GetPortfolio_Call) Return(_a0 *app.Portfolio, _a1 error) *PersonService_GetPortfolio_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PersonService_GetPortfolio_Call) RunAndReturn(run func(context.Context, int64) (*app.Portfolio, error)) *PersonService_GetPortfolio_Call {
	_c.Call.Return(run)
	return _c
}

// NewPersonService creates a new instance of PersonService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPersonService(t interface {
	mock.TestingT
	Cleanup(func())
}) *PersonService {
	mock := &PersonService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
