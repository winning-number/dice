// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dice "github.com/winning-number/dice"
)

// Dice is an autogenerated mock type for the Dice type
type Dice struct {
	mock.Mock
}

type Dice_Expecter struct {
	mock *mock.Mock
}

func (_m *Dice) EXPECT() *Dice_Expecter {
	return &Dice_Expecter{mock: &_m.Mock}
}

// BestOrder provides a mock function with given fields:
func (_m *Dice) BestOrder() []dice.Face {
	ret := _m.Called()

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func() []dice.Face); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_BestOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BestOrder'
type Dice_BestOrder_Call struct {
	*mock.Call
}

// BestOrder is a helper method to define mock.On call
func (_e *Dice_Expecter) BestOrder() *Dice_BestOrder_Call {
	return &Dice_BestOrder_Call{Call: _e.mock.On("BestOrder")}
}

func (_c *Dice_BestOrder_Call) Run(run func()) *Dice_BestOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_BestOrder_Call) Return(faces []dice.Face) *Dice_BestOrder_Call {
	_c.Call.Return(faces)
	return _c
}

func (_c *Dice_BestOrder_Call) RunAndReturn(run func() []dice.Face) *Dice_BestOrder_Call {
	_c.Call.Return(run)
	return _c
}

// Faces provides a mock function with given fields: faces
func (_m *Dice) Faces(faces ...int32) []dice.Face {
	_va := make([]interface{}, len(faces))
	for _i := range faces {
		_va[_i] = faces[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func(...int32) []dice.Face); ok {
		r0 = rf(faces...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_Faces_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Faces'
type Dice_Faces_Call struct {
	*mock.Call
}

// Faces is a helper method to define mock.On call
//   - faces ...int32
func (_e *Dice_Expecter) Faces(faces ...interface{}) *Dice_Faces_Call {
	return &Dice_Faces_Call{Call: _e.mock.On("Faces",
		append([]interface{}{}, faces...)...)}
}

func (_c *Dice_Faces_Call) Run(run func(faces ...int32)) *Dice_Faces_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int32, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(int32)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Dice_Faces_Call) Return(_a0 []dice.Face) *Dice_Faces_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_Faces_Call) RunAndReturn(run func(...int32) []dice.Face) *Dice_Faces_Call {
	_c.Call.Return(run)
	return _c
}

// FacesByNBPick provides a mock function with given fields: _a0
func (_m *Dice) FacesByNBPick(_a0 int64) []dice.Face {
	ret := _m.Called(_a0)

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func(int64) []dice.Face); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_FacesByNBPick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FacesByNBPick'
type Dice_FacesByNBPick_Call struct {
	*mock.Call
}

// FacesByNBPick is a helper method to define mock.On call
//   - _a0 int64
func (_e *Dice_Expecter) FacesByNBPick(_a0 interface{}) *Dice_FacesByNBPick_Call {
	return &Dice_FacesByNBPick_Call{Call: _e.mock.On("FacesByNBPick", _a0)}
}

func (_c *Dice_FacesByNBPick_Call) Run(run func(_a0 int64)) *Dice_FacesByNBPick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *Dice_FacesByNBPick_Call) Return(_a0 []dice.Face) *Dice_FacesByNBPick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_FacesByNBPick_Call) RunAndReturn(run func(int64) []dice.Face) *Dice_FacesByNBPick_Call {
	_c.Call.Return(run)
	return _c
}

// History provides a mock function with given fields:
func (_m *Dice) History() []int32 {
	ret := _m.Called()

	var r0 []int32
	if rf, ok := ret.Get(0).(func() []int32); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	return r0
}

// Dice_History_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'History'
type Dice_History_Call struct {
	*mock.Call
}

// History is a helper method to define mock.On call
func (_e *Dice_Expecter) History() *Dice_History_Call {
	return &Dice_History_Call{Call: _e.mock.On("History")}
}

func (_c *Dice_History_Call) Run(run func()) *Dice_History_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_History_Call) Return(_a0 []int32) *Dice_History_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_History_Call) RunAndReturn(run func() []int32) *Dice_History_Call {
	_c.Call.Return(run)
	return _c
}

// LeastPick provides a mock function with given fields:
func (_m *Dice) LeastPick() []dice.Face {
	ret := _m.Called()

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func() []dice.Face); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_LeastPick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LeastPick'
type Dice_LeastPick_Call struct {
	*mock.Call
}

// LeastPick is a helper method to define mock.On call
func (_e *Dice_Expecter) LeastPick() *Dice_LeastPick_Call {
	return &Dice_LeastPick_Call{Call: _e.mock.On("LeastPick")}
}

func (_c *Dice_LeastPick_Call) Run(run func()) *Dice_LeastPick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_LeastPick_Call) Return(_a0 []dice.Face) *Dice_LeastPick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_LeastPick_Call) RunAndReturn(run func() []dice.Face) *Dice_LeastPick_Call {
	_c.Call.Return(run)
	return _c
}

// MorePick provides a mock function with given fields:
func (_m *Dice) MorePick() []dice.Face {
	ret := _m.Called()

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func() []dice.Face); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_MorePick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MorePick'
type Dice_MorePick_Call struct {
	*mock.Call
}

// MorePick is a helper method to define mock.On call
func (_e *Dice_Expecter) MorePick() *Dice_MorePick_Call {
	return &Dice_MorePick_Call{Call: _e.mock.On("MorePick")}
}

func (_c *Dice_MorePick_Call) Run(run func()) *Dice_MorePick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_MorePick_Call) Return(_a0 []dice.Face) *Dice_MorePick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_MorePick_Call) RunAndReturn(run func() []dice.Face) *Dice_MorePick_Call {
	_c.Call.Return(run)
	return _c
}

// NBPick provides a mock function with given fields: face
func (_m *Dice) NBPick(face int32) int64 {
	ret := _m.Called(face)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int32) int64); ok {
		r0 = rf(face)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Dice_NBPick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NBPick'
type Dice_NBPick_Call struct {
	*mock.Call
}

// NBPick is a helper method to define mock.On call
//   - face int32
func (_e *Dice_Expecter) NBPick(face interface{}) *Dice_NBPick_Call {
	return &Dice_NBPick_Call{Call: _e.mock.On("NBPick", face)}
}

func (_c *Dice_NBPick_Call) Run(run func(face int32)) *Dice_NBPick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32))
	})
	return _c
}

func (_c *Dice_NBPick_Call) Return(_a0 int64) *Dice_NBPick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_NBPick_Call) RunAndReturn(run func(int32) int64) *Dice_NBPick_Call {
	_c.Call.Return(run)
	return _c
}

// NBThrow provides a mock function with given fields:
func (_m *Dice) NBThrow() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Dice_NBThrow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NBThrow'
type Dice_NBThrow_Call struct {
	*mock.Call
}

// NBThrow is a helper method to define mock.On call
func (_e *Dice_Expecter) NBThrow() *Dice_NBThrow_Call {
	return &Dice_NBThrow_Call{Call: _e.mock.On("NBThrow")}
}

func (_c *Dice_NBThrow_Call) Run(run func()) *Dice_NBThrow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_NBThrow_Call) Return(_a0 int64) *Dice_NBThrow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_NBThrow_Call) RunAndReturn(run func() int64) *Dice_NBThrow_Call {
	_c.Call.Return(run)
	return _c
}

// SetThrow provides a mock function with given fields: _a0
func (_m *Dice) SetThrow(_a0 int32) {
	_m.Called(_a0)
}

// Dice_SetThrow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetThrow'
type Dice_SetThrow_Call struct {
	*mock.Call
}

// SetThrow is a helper method to define mock.On call
//   - _a0 int32
func (_e *Dice_Expecter) SetThrow(_a0 interface{}) *Dice_SetThrow_Call {
	return &Dice_SetThrow_Call{Call: _e.mock.On("SetThrow", _a0)}
}

func (_c *Dice_SetThrow_Call) Run(run func(_a0 int32)) *Dice_SetThrow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32))
	})
	return _c
}

func (_c *Dice_SetThrow_Call) Return() *Dice_SetThrow_Call {
	_c.Call.Return()
	return _c
}

func (_c *Dice_SetThrow_Call) RunAndReturn(run func(int32)) *Dice_SetThrow_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *Dice) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Dice_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type Dice_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *Dice_Expecter) String() *Dice_String_Call {
	return &Dice_String_Call{Call: _e.mock.On("String")}
}

func (_c *Dice_String_Call) Run(run func()) *Dice_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_String_Call) Return(_a0 string) *Dice_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_String_Call) RunAndReturn(run func() string) *Dice_String_Call {
	_c.Call.Return(run)
	return _c
}

// Throw provides a mock function with given fields:
func (_m *Dice) Throw() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Dice_Throw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Throw'
type Dice_Throw_Call struct {
	*mock.Call
}

// Throw is a helper method to define mock.On call
func (_e *Dice_Expecter) Throw() *Dice_Throw_Call {
	return &Dice_Throw_Call{Call: _e.mock.On("Throw")}
}

func (_c *Dice_Throw_Call) Run(run func()) *Dice_Throw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_Throw_Call) Return(_a0 int32) *Dice_Throw_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dice_Throw_Call) RunAndReturn(run func() int32) *Dice_Throw_Call {
	_c.Call.Return(run)
	return _c
}

// WeakestOrder provides a mock function with given fields:
func (_m *Dice) WeakestOrder() []dice.Face {
	ret := _m.Called()

	var r0 []dice.Face
	if rf, ok := ret.Get(0).(func() []dice.Face); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dice.Face)
		}
	}

	return r0
}

// Dice_WeakestOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WeakestOrder'
type Dice_WeakestOrder_Call struct {
	*mock.Call
}

// WeakestOrder is a helper method to define mock.On call
func (_e *Dice_Expecter) WeakestOrder() *Dice_WeakestOrder_Call {
	return &Dice_WeakestOrder_Call{Call: _e.mock.On("WeakestOrder")}
}

func (_c *Dice_WeakestOrder_Call) Run(run func()) *Dice_WeakestOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Dice_WeakestOrder_Call) Return(faces []dice.Face) *Dice_WeakestOrder_Call {
	_c.Call.Return(faces)
	return _c
}

func (_c *Dice_WeakestOrder_Call) RunAndReturn(run func() []dice.Face) *Dice_WeakestOrder_Call {
	_c.Call.Return(run)
	return _c
}

// NewDice creates a new instance of Dice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDice(t interface {
	mock.TestingT
	Cleanup(func())
}) *Dice {
	mock := &Dice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}