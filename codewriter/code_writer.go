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

func WritePop(command parser.Command) string {
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

	if command.Arg1 == "constant" {
		// This line depends on constant, local, maybe other things?
		assemblyCode.WriteLine("@" + strconv.Itoa(command.Arg2))
	}

	// Set D to Arg2
	assemblyCode.WriteLine("D=A")

	// Add Arg2 (D) to stack
	assemblyCode.WriteLine("@SP") // SP stores to a memory location
	assemblyCode.WriteLine("A=M") // A points to the memory location of SP
	assemblyCode.WriteLine("M=D") // Set the current memory location (A) to D

	// Increment stack pointer
	assemblyCode.WriteLine("@SP")
	assemblyCode.WriteLine("M=M+1")

	return string(assemblyCode) + "\n"
}

func (ac *AssemblyCode) WriteLine(line string) {
	*ac = *ac + "\n" + AssemblyCode(line)
}
