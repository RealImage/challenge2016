// Package main provides ...
package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributorPermissions(t *testing.T) {
	os.Args = []string{"distributors.txt", "input.txt", "output.txt"}
	expectedOutputFile, err := os.Open("expectedOutput.txt")
	assert.Nil(t, err, "File open error")
	defer expectedOutputFile.Close()

	actualOutputFile, err := os.Open(os.Args[2])
	defer actualOutputFile.Close()

	expectedOutputScanner := bufio.NewScanner(expectedOutputFile)
	actualOutputScanner := bufio.NewScanner(actualOutputFile)
	for actualOutputScanner.Scan() {
		output := actualOutputScanner.Text()
		expectedOutputScanner.Scan()
		expectedOutput := expectedOutputScanner.Text()
		assert.Equal(t, output, expectedOutput, "Output should be the same")
	}
}
