package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if parsedLine, isValid := parser.ParseLine(line); isValid {
			// Turn parsed line into assembly code

			assemblyCode := parsedLine

			// if (command is aritmethic) {
			// assemblyCode := writer.writeArithmetic(parsedLine)
			// }

			// Write assembly code to new file
			// newFile.WriteString(assemblyCode.CommandType)

			fmt.Println(assemblyCode)
		}

	}

}
