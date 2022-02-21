package main

import "fmt"

type InputOutputConfig struct {
	Input struct {
		Stdin bool
		File  string
	}
	Output struct {
		Stdout bool
		File   string
	}
}

func NewInputOutputConfig() *InputOutputConfig {
	return &InputOutputConfig{
		Input: struct {
			Stdin bool
			File  string
		}{
			Stdin: true,
			File:  "",
		},
		Output: struct {
			Stdout bool
			File   string
		}{
			Stdout: true,
			File:   "",
		},
	}
}

func (config *InputOutputConfig) Log() {
	fmt.Printf("config: \n\tOUTPUT:\n\t\tStdout: %v\tFile: %s\n\tINPUT:\n\t\tStdin: %v\tFile: %s\n", config.Output.Stdout, config.Output.File, config.Input.Stdin, config.Input.File)
}
