// Package main provides ...
package main

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributorPermissions(t *testing.T) {
	loadDistributorRules("distributors.txt")
	expectedOutputFile, err := os.Open("expectedOutput.txt")
	assert.Nil(t, err, "File open error")
	defer expectedOutputFile.Close()

	inputFile, err := os.Open("input.txt")
	defer inputFile.Close()

	expectedOutputScanner := bufio.NewScanner(expectedOutputFile)
	inputScanner := bufio.NewScanner(inputFile)
	for inputScanner.Scan() {
		inputTokens := strings.Split(inputScanner.Text(), " ")
		output := computeAndWriteAnswers(inputTokens[0], inputTokens[1])
		expectedOutputScanner.Scan()
		expectedOutput := expectedOutputScanner.Text()
		assert.Equal(t, output, expectedOutput, "Output not same for: "+inputTokens[0]+" "+inputTokens[1])
	}
}
