## go-test

A lightweight dependency-free helper for writing golang tests.

[![codecov](https://codecov.io/gh/pjbgf/go-test/branch/master/graph/badge.svg)](https://codecov.io/gh/pjbgf/go-test)
[![GoReport](https://goreportcard.com/badge/github.com/pjbgf/go-test)](https://goreportcard.com/report/github.com/pjbgf/go-test)
[![GoDoc](https://godoc.org/github.com/pjbgf/go-test?status.svg)](https://godoc.org/github.com/pjbgf/go-test)
![build](https://github.com/pjbgf/go-test/workflows/go/badge.svg)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](http://choosealicense.com/licenses/mit/)



Sample Code
```golang
package calc

func Sum(value1, value2 int) int {
	return value1 + value2
}
```

Test for Sample Code
```golang
package calc

import (
	"testing"

	"github.com/pjbgf/go-test/should"
)

func TestSum(t *testing.T) {
	assertThat := func(assumption string, value1, value2, expected int) {
		should := should.New(t)

		actual := Sum(value1, value2)

		should.BeEqual(expected, actual, assumption)
	}

	assertThat("should return 13 for 4 and 9", 4, 9, 13)
	assertThat("should return 50 for 15 and 30", 15, 35, 50)
}
```