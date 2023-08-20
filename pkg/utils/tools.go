package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

// Set the SSH environment variable to "~/.ssh/"
var _ = os.Setenv("SSH", "~/.ssh/")

var (
	// This channel receives the output of the `ExecuteWithInventory` method of the `Playbook` object line by line.
	// If you want to analyze or interpret the output, you can listen to it after executing this method `ExecuteWithInventory`.
	ExecutionWithInventoryOutputPipeline = make(chan string, 1000)

	// This channel is used internally.
	pipeline = make(chan string)
)

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) (string, error) {
	// Generate a UUID (Universally Unique Identifier) and convert it to a string.
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	randomString := uuidObj.String()
	// Truncate the string to the desired length.
	if len(randomString) > length {
		randomString = randomString[:length]
	}

	return randomString, nil
}

// GetTheOutput function will retrieve each line of the output and provide it to the output pipeline.
func GetTheOutput(Stdout io.ReadCloser) {
	scanner := bufio.NewScanner(Stdout)
	for scanner.Scan() {
		pipeline <- scanner.Text()
		ExecutionWithInventoryOutputPipeline <- scanner.Text()
	}
	close(pipeline)
	close(ExecutionWithInventoryOutputPipeline)
}

// PrintOutputs function prints the output received from the pipeline.
func PrintOutputs() {
	for theOutput := range pipeline {
		fmt.Println(theOutput)
	}
}