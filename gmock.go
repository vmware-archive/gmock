package gmock

import (
	"reflect"
)

type GMock struct {
	target   reflect.Value
	original reflect.Value
}

func MockTarget(targetVar interface{}) *GMock {
	mock := &GMock{}
	mock.target = reflect.ValueOf(targetVar).Elem()
	mock.target = reflect.New(mock.target.Type()).Elem()
	mock.original.Set(mock.target)
	return mock
}

func MockTargetWithValue(targetVar interface{}, mockValue interface{}) *GMock {
	mock := MockTarget(targetVar)
	mock.Replace(mockValue)
	return mock
}

func (self *GMock) Replace(mockValue interface{}) {
	replacement := reflect.ValueOf(mockValue)

	if !replacement.IsValid() {
		replacement = reflect.Zero(self.target.Type())
	}

	self.target.Set(replacement)
}

func (self *GMock) Restore() {
	self.target.Set(self.original)
}

func (self *GMock) GetTarget() reflect.Value {
	return self.target
}

func (self *GMock) GetOriginal() reflect.Value {
	return self.original
}
