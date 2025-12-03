package helpers

import (
	"fmt"
	"log"
)

type Testcase struct {
	FileName       string
	ExpectedResult int
}

func LogLine(a ...any) {
	fmt.Println(a...)
}

type Solver func(fileName string, first bool) int

func RunTests(tests []Testcase, method Solver, versionIndicator bool) int {
	errors := 0

	for index := range tests {
		testcase := tests[index]
		result := method(testcase.FileName, versionIndicator)
		if result != testcase.ExpectedResult {
			log.Fatal("Test failed", testcase, result)
			errors += 1
		}

		fmt.Println("Test succeeded", testcase.FileName)
	}

	return errors
}
