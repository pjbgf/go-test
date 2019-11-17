package should

import (
	"errors"
	"fmt"
	"testing"
)

type testingStub struct {
	helperCalled bool
	logMessage   string
	errorMessage string
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

func (t *testingStub) Error(args ...interface{}) {
	t.errorMessage = fmt.Sprint(args...)
}

func TestBeNil(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}, expectedErrorMessage string) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)

			should.BeNil(value, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail for non empty string",
			"test",
			"value was expected to be nil, but was 'test' instead")
		assertThat("should fail for empty string",
			"",
			"value was expected to be nil, but was '' instead")
		assertThat("should fail for non-nil func()",
			func() {},
			"value was expected to be nil, but was an initialised value of type func() instead")
		assertThat("should fail for non-nil chan int",
			make(chan int),
			"value was expected to be nil, but was an initialised value of type chan int instead")
		assertThat("should fail for non-nil map[int]int",
			make(map[int]int),
			"value was expected to be nil, but was an initialised value of type map[int]int instead")
		assertThat("should fail for non-nil []uint8",
			make([]uint8, 0),
			"value was expected to be nil, but was an initialised value of type []uint8 instead")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeNil(value, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
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
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)
			expectedErrorMessage := "value was expected to not be nil, but it was not"

			should.BeNotNil(value, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
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

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
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
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)
			expectedErrorMessage := "error was expected, but did not happen"

			should.Error(err, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail for nil", nil)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, err error) {
			stub := testingStub{}
			should := New(&stub)

			should.Error(err, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
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
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)
			expectedErrorMessage := "error was not expected, but did happen"

			should.NotError(err, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
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

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
			}
		}

		assertThat("should not fail for nil", nil)
	})
}

func TestBeEqual(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}, expectedErrorMessage string) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)

			should.BeEqual(expected, actual, assumption)

			if !stub.WasHelperCalled() {
				t.Error("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail diff strings and escape errors with line breaks",
			"ab\nc", "cde\n",
			"expected 'ab\\nc' but got 'cde\\n' instead")
		assertThat("should fail diff strings and escape errors with tabs",
			"ab\tc", "cde\t",
			"expected 'ab\\tc' but got 'cde\\t' instead")
		assertThat("should fail for true and \"true\"",
			true, "true",
			"expected '%!s(bool=true)' but got 'true' instead")
		assertThat("should fail for (int32=6) and (int16=6)",
			int32(6), int16(6),
			"expected '%!s(int32=6)' but got '%!s(int16=6)' instead")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeEqual(expected, actual, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
			}
		}

		assertThat("should not fail tests for equal strings", "abc", "abc")
		assertThat("should not fail tests for equal boolean", true, true)
		assertThat("should not fail tests for equal []byte and []uint8", []byte{2}, []uint8{2})
	})
}

func TestBeNotEqual(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}, expectedErrorMessage string) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)

			should.BeNotEqual(expected, actual, assumption)

			if !stub.WasHelperCalled() {
				t.Error("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail tests for equal strings",
			"abc", "abc",
			"expected 'abc' not to be equal to 'abc', but they were")
		assertThat("should fail tests for equal boolean",
			true, true,
			"expected '%!s(bool=true)' not to be equal to '%!s(bool=true)', but they were")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, expected, actual interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.BeNotEqual(expected, actual, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
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
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)
			expectedErrorMessage := "value was expected to be true, but was false instead"

			should.BeTrue(value, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail for false", false)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)

			should.BeTrue(value, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
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
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)
			expectedErrorMessage := "value was expected to be false, but was true instead"

			should.BeFalse(value, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail for true", true)
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value bool) {
			stub := testingStub{}
			should := New(&stub)

			should.BeFalse(value, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
			}
		}

		assertThat("should not fail tests for false", false)
	})
}

func TestHaveSameType(t *testing.T) {
	t.Run("scenarios that must fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value1, value2 interface{}, expectedErrorMessage string) {
			stub := testingStub{}
			should := New(&stub)
			expectedLogMessage := fmt.Sprintf("==== %s ====", assumption)

			should.HaveSameType(value1, value2, assumption)

			if !stub.WasHelperCalled() {
				t.Errorf("Helper() call was expected but did not happen")
			}
			if expectedLogMessage != stub.logMessage {
				t.Errorf("wanted '%s' got '%s'", expectedLogMessage, stub.logMessage)
			}
			if expectedErrorMessage != stub.errorMessage {
				t.Errorf("wanted '%s' got '%s'", expectedErrorMessage, stub.errorMessage)
			}
		}

		assertThat("should fail for true and 'true'", true, "true", "types expected to be same, but were 'bool' and 'string'")
		assertThat("should fail for \"a\" and ''", "a", 'a', "types expected to be same, but were 'string' and 'int32'")
	})

	t.Run("scenarios that must not fail tests", func(t *testing.T) {
		assertThat := func(assumption string, value1, value2 interface{}) {
			stub := testingStub{}
			should := New(&stub)

			should.HaveSameType(value1, value2, assumption)

			if stub.WasHelperCalled() {
				t.Error("Helper() call was not expected")
			}
			if stub.logMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.logMessage)
			}
			if stub.errorMessage != "" {
				t.Errorf("wanted '%s' got '%s'", "", stub.errorMessage)
			}
		}

		assertThat("should not fail for true and false", true, false)
		assertThat("should not fail for 'something' and 'blah'", "something", "blah")
		assertThat("should not fail for []uint8{1} and []byte{1}", []uint8{1}, []byte{1})
	})
}
