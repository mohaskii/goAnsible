package utils

import (
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

// Playbook represents an Ansible playbook.
type Playbook struct {
	TheConfigs []ConFig
}

// ConFig represents a configuration within an Ansible playbook.
type ConFig struct {
	Name       string        `yaml:"name"`
	Hosts      string        `yaml:"hosts"`
	RemoteUser string        `yaml:"remote_user"`
	Tasks      []interface{} `yaml:"tasks"`
}

// ExecuteWithInventory executes the playbook with the specified inventory file and optional flags.
func (p *Playbook) ExecuteWithInventory(inventoryName string, flags ...string) (err error) {
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
	//*****************************************
	cmd := exec.Command("ansible-playbook", "-i", inventoryName, tempFile)
	cmd.Args = append(cmd.Args, flags...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	go PrintOutputs()
	go GetTheOutput(stdout)
	if err = cmd.Start(); err != nil {
		return err
	}
	cmd.Wait()
	os.Remove(tempFile)
	return nil
}

// ConvertToYamlFile converts the playbook configuration to a YAML file with the specified name.
func (p *Playbook) ConvertToYamlFile(fileName string) error {

	out, err := yaml.Marshal(p.TheConfigs)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, out, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil
}