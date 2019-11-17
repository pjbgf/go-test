// Package should provide methods for testing go applications.
package should

import (
	"fmt"
	"reflect"
	"strings"
)

// Should define easy to use methods for testing go applications.
type Should struct {
	t testingT
}

type testingT interface {
	Helper()
	Log(args ...interface{})
	Error(args ...interface{})
}

// New initialises a new Should instance.
func New(t testingT) *Should {
	return &Should{t}
}

// BeNil fails the test if value is not nil.
func (s *Should) BeNil(value interface{}, assumption string) {
	if !isNil(value) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error(fmt.Sprintf("value was expected to be nil, but was %s instead", getNotNilString(value)))
	}
}

// BeNotNil fails the test if value is nil.
func (s *Should) BeNotNil(value interface{}, assumption string) {
	if isNil(value) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error("value was expected to not be nil, but it was not")
	}
}

// Error fails the test if err is nil.
func (s *Should) Error(err error, assumption string) {
	if isNil(err) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error("error was expected, but did not happen")
	}
}

// NotError fails the test if err is not nil.
func (s *Should) NotError(err error, assumption string) {
	if !isNil(err) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error("error was not expected, but did happen")
	}
}

// BeEqual compares the values of both expected and actual and fails the test if they differ.
func (s *Should) BeEqual(expected, actual interface{}, assumption string) {
	if !reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error(fmt.Sprintf("expected '%s' but got '%s' instead", escape(expected), escape(actual)))
	}
}

// BeNotEqual compares the values of both expected and actual and fails the test if they are equal.
func (s *Should) BeNotEqual(expected, actual interface{}, assumption string) {
	if reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error(fmt.Sprintf("expected '%s' not to be equal to '%s', but they were", escape(expected), escape(actual)))
	}
}

// BeTrue fails the test if value is false.
func (s *Should) BeTrue(value bool, assumption string) {
	if !value {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error("value was expected to be true, but was false instead")
	}
}

// BeFalse fails the test if value is true.
func (s *Should) BeFalse(value bool, assumption string) {
	if value {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error("value was expected to be false, but was true instead")
	}
}

// HaveSameType compares the types of both expected and actual and fails the test if they differ.
func (s *Should) HaveSameType(expected, actual interface{}, assumption string) {
	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		s.t.Helper()
		s.t.Log(fmt.Sprintf("==== %s ====", assumption))
		s.t.Error(fmt.Sprintf("types expected to be same, but were '%s' and '%s'", expectedType, actualType))
	}
}

func escape(value interface{}) interface{} {
	tmpValue, ok := value.(string)
	if ok {
		tmpValue = strings.ReplaceAll(tmpValue, "\n", "\\n")
		return strings.ReplaceAll(tmpValue, "\t", "\\t")
	}

	return value
}

func isNil(value interface{}) bool {
	if value == nil {
		return true
	}

	t := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map,
		reflect.Func, reflect.Chan:
		return t.IsNil()
	}

	return false
}

func getNotNilString(value interface{}) string {
	t := reflect.ValueOf(value)
	switch t.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map,
		reflect.Func, reflect.Chan:
		return fmt.Sprintf("an initialised value of type %s", t.Type())
	}

	return fmt.Sprintf("'%s'", value)
}
