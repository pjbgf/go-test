package should

import (
	"errors"
	"fmt"
	"testing"
)

type testingStub struct {
	hasFailed    bool
	helperCalled bool
	logMessage   string
}

func (t *testingStub) Helper() {
	t.helperCalled = true
}

func (t *testingStub) WasHelperCalled() bool {
	return t.helperCalled
}

func (t *testingStub) Log(args ...interface{}) {
	t.logMessage = fmt.Sprint(args...)
}

func (t *testingStub) Fail() {
	t.hasFailed = true
}

func TestBeNil(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: '%s'",
				assumption, "BeNil", "nil", value)

			should.BeNil(value, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for non empty string", "test")
		assertThat("should fail for empty string", "")
		assertThat("should fail for non-nil func()", func() {})
		assertThat("should fail for non-nil chan int", make(chan int))
		assertThat("should fail for non-nil map[int]int", make(map[int]int))
		assertThat("should fail for non-nil []uint8", make([]uint8, 0))
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeNil(value, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for nil", nil)
		assertThat("should not fail tests for interface{}(nil)", interface{}(nil))
		assertThat("should not fail tests for (*int)(nil)", (*int)(nil))
		assertThat("should not fail tests for ([]byte)(nil)", ([]byte)(nil))
		assertThat("should not fail tests for (map[bool]bool)(nil)", (map[bool]bool)(nil))
		assertThat("should not fail tests for (func())(nil)", (func())(nil))
		assertThat("should not fail tests for (chan int)(nil)", (chan int)(nil))
	})
}

func TestNotBeNil(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: '%s'",
				assumption, "BeNotNil", "!= nil", value)

			should.BeNotNil(value, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for untyped nil", nil)
		assertThat("should fail for interface{}(nil)", interface{}(nil))
		assertThat("should fail for (*int)(nil)", (*int)(nil))
		assertThat("should fail for (func())(nil))", (func())(nil))
		assertThat("should fail for (chan int)(nil))", (chan int)(nil))
		assertThat("should fail for nil (map[bool]bool)(nil))", (map[bool]bool)(nil))
		assertThat("should fail for nil ([]byte)(nil)", ([]byte)(nil))
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeNotNil(value, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for empty string", "test")
		assertThat("should not fail tests for empty string", "")
		assertThat("should not fail tests for non-nil func()", func() {})
		assertThat("should not fail tests for non-nil chan int", make(chan int))
		assertThat("should not fail tests for non-nil map[int]int", make(map[int]int))
		assertThat("should not fail tests for non-nil []uint8", make([]uint8, 0))
	})
}

func TestError(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, err error) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: '%s'",
				assumption, "Error", "!= nil", err)

			should.Error(err, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for nil", nil)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, err error) {
			stub := testingStub{}
			should := New(&stub)

			should.Error(err, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for empty error", errors.New(""))
		assertThat("should not fail tests for empty error", errors.New("some error"))
	})
}

func TestNotError(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, err error) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: '%s'",
				assumption, "NotError", "nil", err)

			should.NotError(err, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail tests for empty error", errors.New(""))
		assertThat("should fail tests for non empty error", errors.New("some error"))
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, err error) {
			stub := testingStub{}
			should := New(&stub)

			should.NotError(err, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail for nil", nil)
	})
}

func TestBeEqual(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: '%s'\n    actual: '%s'",
				assumption, "BeEqual", escape(expected), escape(actual))

			should.BeEqual(expected, actual, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Error("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail diff strings and escape errors with line breaks",
			"ab\nc", "cde\n")
		assertThat("should fail diff strings and escape errors with tabs",
			"ab\tc", "cde\t")
		assertThat("should fail for true and \"true\"",
			true, "true")
		assertThat("should fail for (int32=6) and (int16=6)",
			int32(6), int16(6))
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeEqual(expected, actual, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for equal strings", "abc", "abc")
		assertThat("should not fail tests for equal strings", "abc\n", `abc
`)
		assertThat("should not fail tests for equal boolean", true, true)
		assertThat("should not fail tests for equal []byte and []uint8", []byte{2}, []uint8{2})
	})
}

func TestBeNotEqual(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: '%s'\n    actual: '%s'",
				assumption, "BeNotEqual", escape(expected), escape(actual))

			should.BeNotEqual(expected, actual, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Error("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail tests for equal strings", "abc", "abc")
		assertThat("should fail tests for equal boolean", true, true)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeNotEqual(expected, actual, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for different strings", "abc", "def")
		assertThat("should not fail tests for different booleans", true, false)
	})
}

func TestBeTrue(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: %t",
				assumption, "BeTrue", "true", value)

			should.BeTrue(value, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for false", false)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)

			should.BeTrue(value, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for true", true)
	})
}

func TestBeFalse(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: %t",
				assumption, "BeFalse", "false", value)

			should.BeFalse(value, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for true", true)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)

			should.BeFalse(value, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail tests for false", false)
	})
}

func TestHaveSameType(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value1, value2 interface{}, type1, type2 string) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("\nassumption: [ %s ]\n    should: %s \n  expected: %s\n    actual: %s",
				assumption, "HaveSameType", type1, type2)

			should.HaveSameType(value1, value2, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for true and 'true'", true, "true", "bool", "string")
		assertThat("should fail for \"a\" and ''", "a", 'a', "string", "int32")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value1, value2 interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.HaveSameType(value1, value2, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should not fail for true and false", true, false)
		assertThat("should not fail for 'something' and 'blah'", "something", "blah")
		assertThat("should not fail for []uint8{1} and []byte{1}", []uint8{1}, []byte{1})
	})
}

func TestHaveSameItems(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{},
			expectedLogMessage string) {
			stub := testingStub{}
			should := New(&stub)

			should.HaveSameItems(expected, actual, assumption)

			if !stub.hasFailed {
				t.Error("test was expected to fail but did not")
			}
			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
		}

		assertThat("should fail for missing item in []string",
			[]string{"a", "b", "c"}, []string{"c", "b", "d"},
			"\nassumption: [ should fail for missing item in []string ]\n    should: HaveSameItems \n    reason: items missing\n  expected: [a b c]\n    actual: [c b d]\n   missing: [a]")
		assertThat("should fail for missing item in []int",
			[]int{5, 7, 2}, []int{2, 5, 1},
			"\nassumption: [ should fail for missing item in []int ]\n    should: HaveSameItems \n    reason: items missing\n  expected: [5 7 2]\n    actual: [2 5 1]\n   missing: [7]")
		assertThat("should fail for different types",
			[]int{5, 7, 2}, []string{"5", "7", "2"},
			"\nassumption: [ should fail for different types ]\n    should: HaveSameItems \n    reason: type mismatch\n  expected: []int\n    actual: []string")
		assertThat("should fail for different lengths",
			[]string{"a", "b", "c"}, []string{"c", "b", "d", "a"},
			"\nassumption: [ should fail for different lengths ]\n    should: HaveSameItems \n    reason: length mismatch\n  expected: [a b c]\n    actual: [c b d a]\nlength exp: 3\nlength act: 4")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.HaveSameItems(expected, actual, assumption)

			if stub.hasFailed {
				t.Error("test was expected to not fail but it did")
			}
			if stub.WasHelperCalled() {
				t.Errorf("Helper() call was not expected but it did happen")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
		}

		assertThat("should fail for missing item in []string",
			[]string{"a", "b", "c"}, []string{"c", "b", "a"})
		assertThat("should fail for missing item in []int",
			[]int{5, 7, 2}, []int{2, 5, 7})
	})
}
