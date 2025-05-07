package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"nandtotetris/vmtranslator/codewriter"
	"nandtotetris/vmtranslator/parser"
)

func main() {
	fileName := os.Args[1]

	// Open the file to read from
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	// Create file to write to
	newFileName := fileName[:strings.Index(fileName, ".")]
	newFile, err := os.Create(newFileName + ".asm")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer newFile.Close()

	// Init symbols
	// These are initialised by a test script in the cpu emulator. Need to
	// define them here to run in the emulator without a test script.
	newFile.WriteString(codewriter.WriteVariable("SP", 256))
	newFile.WriteString(codewriter.WriteVariable("LCL", 300))
	newFile.WriteString(codewriter.WriteVariable("ARG", 400))
	newFile.WriteString(codewriter.WriteVariable("THIS", 3000))
	newFile.WriteString(codewriter.WriteVariable("THAT", 3010))

	scanner := bufio.NewScanner(file)
	count := 1

	for scanner.Scan() {
		line := scanner.Text()

		if parsedLine, isValid := parser.ParseLine(line); isValid {
			// Write the line as a comment to explain what the following instructions do
			newFile.WriteString("\n// VM line: " + line)

			code := ""
			if parsedLine.CommandType == parser.PUSH_COMMAND {
				code = codewriter.WritePush(parsedLine, newFileName)
			}

			if parsedLine.CommandType == parser.POP_COMMAND {
				code = codewriter.WritePop(parsedLine, newFileName)
			}

			if parsedLine.CommandType == parser.ARIT_COMMAND {
				code = codewriter.WriteArithmetic(parsedLine, count)
			}

			newFile.WriteString(code)
			// fmt.Println(parsedLine)
		}

		count++
	}

	// End with intifite loop
	newFile.WriteString("\n(END)\n@END\n0;JMP")
}
