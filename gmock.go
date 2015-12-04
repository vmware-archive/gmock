package gmock

import (
	"reflect"
)

type GMock struct {
	Target   reflect.Value
	Original reflect.Value
}

func MockTarget(targetVar interface{}) *GMock {
	mock := &GMock{}
	mock.Target = reflect.ValueOf(targetVar).Elem()
	mock.Target = reflect.New(mock.Target.Type()).Elem()
	mock.Original.Set(mock.Target)
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
		replacement = reflect.Zero(self.Target.Type())
	}

	self.Target.Set(replacement)
}

func (self *GMock) Restore() {
	self.Target.Set(self.Original)
}
