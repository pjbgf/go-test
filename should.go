// Package should provide methods for testing go applications.
package should

import (
	"reflect"
	"testing"
)

// Should define easy to use methods for testing go applications.
type Should struct {
	t *testing.T
}

// New initialises a new Should instance.
func New(t *testing.T) *Should {
	return &Should{t}
}

// BeNil fails the test if value is not nil.
func (s *Should) BeNil(value interface{}, assumption string) {
	if value != nil {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Errorf("value was expected to be nil, but was '%s' instead", value)
	}
}

// BeNotNil fails the test if value is nil.
func (s *Should) BeNotNil(value interface{}, assumption string) {
	if value == nil {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Error("value was not expected to be nil, but it was")
	}
}

// Error fails the test if err is nil.
func (s *Should) Error(err error, assumption string) {
	if err == nil {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Error("error was expected, but did not happen")
	}
}

// NotError fails the test if err is not nil.
func (s *Should) NotError(err error, assumption string) {
	if err != nil {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Error("error was not expected, but did happen")
	}
}

// BeEqual compares the values of both expected and actual and fails the test if they differ.
func (s *Should) BeEqual(expected, actual interface{}, assumption string) {
	if !reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Errorf("expected '%s' but got '%s' instead", expected, actual)
	}
}

// BeNotEqual compares the values of both expected and actual and fails the test if they are equal.
func (s *Should) BeNotEqual(expected, actual interface{}, assumption string) {
	if reflect.DeepEqual(expected, actual) {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Errorf("expected '%s' not to be equal to '%s', but they were", expected, actual)
	}
}

// BeTrue fails the test if value is false.
func (s *Should) BeTrue(value bool, assumption string) {
	if !value {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Error("value was expected to be true, but was false instead")
	}
}

// BeFalse fails the test if value is true.
func (s *Should) BeFalse(value bool, assumption string) {
	if value {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Error("value was expected to be false, but was true instead")
	}
}

// HaveSameType compares the types of both expected and actual and fails the test if they differ.
func (s *Should) HaveSameType(expected, actual interface{}, assumption string) {
	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		s.t.Helper()
		s.t.Log(assumption)
		s.t.Errorf("expected type was '%s' but got '%s' instead", expectedType, actualType)
	}
}
