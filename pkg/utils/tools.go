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
func getTheOutput(Stdout io.ReadCloser, p playbook) {
	scanner := bufio.NewScanner(Stdout)
	for scanner.Scan() {
		p.pipeline <- scanner.Text()
		p.ExecutionWithInventoryOutputPipeline <- scanner.Text()
	}
	close(p.pipeline)
	close(p.ExecutionWithInventoryOutputPipeline)
}

// PrintOutputs function prints the output received from the pipeline.
func printOutputs(p playbook) {
	for theOutput := range p.pipeline {
		fmt.Println(theOutput)
	}
}