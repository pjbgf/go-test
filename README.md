## should

A lightweight dependency-free helper for writing golang tests.

[![codecov](https://codecov.io/gh/pjbgf/should/branch/master/graph/badge.svg)](https://codecov.io/gh/pjbgf/should)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=pjbgf/should)](https://dependabot.com)
[![GoReport](https://goreportcard.com/badge/github.com/pjbgf/should)](https://goreportcard.com/report/github.com/pjbgf/should)
[![GoDoc](https://godoc.org/github.com/pjbgf/should?status.svg)](https://godoc.org/github.com/pjbgf/should)
![build](https://github.com/pjbgf/should/workflows/go/badge.svg)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](http://choosealicense.com/licenses/mit/)



Code to be tested
```golang
func getArchitectures(targetArchitectures []string) []specs.Arch {
	arches := make([]specs.Arch, 0, 20)
	for _, arch := range targetArchitectures {
		switch arch {
		case "amd64":
			arches = append(arches, specs.ArchX86_64, specs.ArchX86, specs.ArchX32)
		case "arm64":
			arches = append(arches, specs.ArchARM, specs.ArchAARCH64)
		}
	}

	return arches
}
```

Test Code
```golang
func TestGetArchitectures(t *testing.T) {
	assertThat := func(assumption string, targetArchitectures []string, expected []specs.Arch) {
		should := should.New(t)
		actual := getArchitectures(targetArchitectures)

		should.BeEqual(expected, actual, assumption)
	}

    assertThat("should return empty archs for no target architectures",
		[]string{},
		[]specs.Arch{})
	assertThat("should support amd64",
		[]string{"amd64"},
		[]specs.Arch{specs.ArchX86_64, specs.ArchX86, specs.ArchX32})
	assertThat("should support arm64",
		[]string{"arm64"},
		[]specs.Arch{specs.ArchARM, specs.ArchAARCH64})
	assertThat("should combine multiple architectures",
		[]string{"amd64", "arm64"},
		[]specs.Arch{specs.ArchX86_64, specs.ArchX86, specs.ArchX32, specs.ArchARM, specs.ArchAARCH64})
}
```