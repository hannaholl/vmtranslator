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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if parsedLine, isValid := parser.ParseLine(line); isValid {
			// Write the line as a comment to explain what the following instructions do
			newFile.WriteString("// VM line: " + line)

			if parsedLine.CommandType == parser.PUSH_COMMAND {
				code := codewriter.WritePush((parsedLine))
				newFile.WriteString(code)
			}

			fmt.Println(parsedLine)
		}
	}

	// End with intifite loop
	newFile.WriteString("(END)\n@END\n0;JMP")
}
