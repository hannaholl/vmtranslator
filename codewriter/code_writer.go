package codewriter

import (
	"nandtotetris/vmtranslator/parser"
	"strconv"
)

// func WriteArithmetic(command parser.Command) string {

// }

// func WritePushPop(command parser.Command) string {

// }

type AssemblyCode string

func WritePush(command parser.Command) string {
	/*
		// push constant 10
		@10
		D=A

		@SP
		A=M
		M=D

		@SP
		M=M+1
	*/

	var assemblyCode AssemblyCode = ""

	// Could write directly to the file instead of a variable?

	switch command.Arg1 {
	case "constant":
		assemblyCode.WriteLine("@" + strconv.Itoa(command.Arg2))
		assemblyCode.WriteLine("D=A") // Set D to Arg2
	case "local":
		assemblyCode.WriteLine("@LCL")
		assemblyCode.WriteLine("D=M")                            // D is the base of local vars
		assemblyCode.WriteLine("@" + strconv.Itoa(command.Arg2)) // index from command
		assemblyCode.WriteLine("A=D+A")                          // load the local address from the command
		assemblyCode.WriteLine("D=M")                            // set D to the local value
	}

	// Add Arg2 (D) to stack
	assemblyCode.WriteLine("@SP") // SP stores a memory location
	assemblyCode.WriteLine("A=M") // A points to the memory location from SP
	assemblyCode.WriteLine("M=D") // Set the current memory location (A) to D

	// Increment stack pointer
	assemblyCode.WriteLine("@SP")
	assemblyCode.WriteLine("M=M+1")

	return string(assemblyCode) + "\n"
}

func (ac *AssemblyCode) WriteLine(line string) {
	*ac = *ac + "\n" + AssemblyCode(line)
}
