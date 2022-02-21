package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var config *InputOutputConfig = NewInputOutputConfig()

	outputFileFlag := flag.String("o", "", "output file")
	inputFileFlag := flag.String("i",
		"", "source code file")
	flag.Parse()

	if *outputFileFlag != "" {
		config.Output.Stdout = false
		config.Output.File = *outputFileFlag
	}
	if *inputFileFlag != "" {
		config.Input.Stdin = false
		config.Input.File = *inputFileFlag
	}

	var lexer *Lexer

	if config.Input.Stdin {
		r := bufio.NewReader(os.Stdin)
		bx, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		lexer = NewLexer(bytes.NewReader(bx))
	} else {
		f, err := os.Open(config.Input.File)
		if err != nil {
			// panic, for now
			panic(err)
		}

		lexer = NewLexer(f)
	}

	output, err := lexer.JSON()
	if err != nil {
		panic(err)
	}

	if config.Output.File != "" {
		err = os.WriteFile(config.Output.File, output, 0644)
		if err != nil {
			log.Printf("[ERR] writing output to file: %s\n", err.Error())
			return
		}
		return
	}

	fmt.Println(string(output))
}
