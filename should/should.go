// Package should provide methods for testing go applications.
package should

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	stringLogFormat        string = "\nassumption: [ %s ]\n    should: %s \n  expected: '%s'\n    actual: '%s'"
	singleValueLogFormat   string = "\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: '%s'"
	singleBooleanLogFormat string = "\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: %t"
	typeLogFormat          string = "\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: %s"
)

// Should define easy to use methods for testing go applications.
type Should struct {
	t testingT
}

type testingT interface {
	Helper()
	Log(args ...interface{})
	Fail()
}

// New initialises a new Should instance.
func New(t testingT) *Should {
	return &Should{t}
}

// BeNil fails the test if value is not nil.
func (s *Should) BeNil(value interface{}, assumption string) {
	if !isNil(value) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleValueLogFormat, assumption, "BeNil", "nil", value))
		s.t.Fail()
	}
}

// BeNotNil fails the test if value is nil.
func (s *Should) BeNotNil(value interface{}, assumption string) {
	if isNil(value) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleValueLogFormat, assumption, "BeNotNil", "!= nil", value))
		s.t.Fail()
	}
}

// Error fails the test if err is nil.
func (s *Should) Error(err error, assumption string) {
	if isNil(err) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleValueLogFormat, assumption, "Error", "!= nil", err))
		s.t.Fail()
	}
}

// NotError fails the test if err is not nil.
func (s *Should) NotError(err error, assumption string) {
	if !isNil(err) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleValueLogFormat, assumption, "NotError", "nil", err))
		s.t.Fail()
	}
}

// BeEqual compares the values of both expected and actual and fails the test if they differ.
func (s *Should) BeEqual(expected, actual interface{}, assumption string) {
	if !reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(stringLogFormat, assumption, "BeEqual",
			escape(expected), escape(actual)))
		s.t.Fail()
	}
}

// BeNotEqual compares the values of both expected and actual and fails the test if they are equal.
func (s *Should) BeNotEqual(expected, actual interface{}, assumption string) {
	if reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(stringLogFormat, assumption, "BeNotEqual",
			escape(expected), escape(actual)))
		s.t.Fail()
	}
}

// BeTrue fails the test if value is false.
func (s *Should) BeTrue(value bool, assumption string) {
	if !value {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleBooleanLogFormat, assumption, "BeTrue", "true", value))
		s.t.Fail()
	}
}

// BeFalse fails the test if value is true.
func (s *Should) BeFalse(value bool, assumption string) {
	if value {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(singleBooleanLogFormat, assumption, "BeFalse", "false", value))
		s.t.Fail()
	}
}

// HaveSameType compares the types of both expected and actual and fails the test if they differ.
func (s *Should) HaveSameType(expected, actual interface{}, assumption string) {
	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		s.t.Helper()
		s.t.Log(fmt.Sprintf(typeLogFormat, assumption, "HaveSameType", expectedType, actualType))
		s.t.Fail()
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
