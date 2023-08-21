package utils

import (
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Playbook represents an Ansible playbook.
type playbook struct {
	// Configs represents a configuration within an Ansible playbook.
	Configs []interface{}
	// ExecutionWithInventoryOutputPipeline receives the output of the `ExecuteWithInventory` method of the `Playbook` object line by line.
	// You can listen to it if you want to analyze or interpret the output after executing the `ExecuteWithInventory` method.
	ExecutionWithInventoryOutputPipeline chan string
	// This channel is used internally to print the output.
	pipeline chan string
	// Set the variable to true if you want to hide the output.
	HideOutput bool
}

// ExecuteWithInventory executes the playbook with the specified inventory file and optional flags.
func (p *playbook) ExecuteWithInventory(inventoryName string, flags ...string) (err error) {
	// Create the YAML file of the playbook.
	tempFile, err := GenerateRandomString(5)
	if err != nil {
		return err
	}
	tempFile = "." + tempFile + ".yml"
	err = p.ConvertToYamlFile(tempFile)
	if err != nil {
		return err
	}
	cmd := exec.Command("ansible-playbook", "-i", inventoryName, tempFile)
	cmd.Args = append(cmd.Args, flags...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if !p.HideOutput {
		go printOutputs(*p)
	}
	go getTheOutput(stdout, *p)
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	os.Remove(tempFile)
	return nil
}

// ConvertToYamlFile converts the playbook configuration to a YAML file with the specified name.
func (p *playbook) ConvertToYamlFile(fileName string) error {

	out, err := yaml.Marshal(p.Configs)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, out, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}

// InitPlaybook initializes a new playbook instance with the specified length of the output buffer.
func InitPlaybook(LenOfTheOutputBuffer int) playbook {
	return playbook{
		Configs:                              make([]interface{}, 0),
		ExecutionWithInventoryOutputPipeline: make(chan string, LenOfTheOutputBuffer),
		pipeline:                             make(chan string),
	}
}
