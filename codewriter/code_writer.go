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
		// pop local 4

		// update LCL to the new address
		@4
		D=A
		@LCL
		M=D+M

		// point A to the location pop from
		@SP
		A=M-1
		//D=M
		//A=D-1

		// the value from the stack
		D=M

		// write value to current local
		@LCL
		M=D

		// decrement SP
		@SP
		M=M-1

		// reset LCL
		@LCL
		D=M
		@4
		M=D-A
	*/
	var assemblyCode AssemblyCode = ""

	// TODO: Support more segments later
	if command.Arg1 == "local" {
		assemblyCode.WriteLine("// update LCL to the new address")
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("D=A")
		assemblyCode.WriteLine("@LCL")
		assemblyCode.WriteLine("M=D+M")
	}

	assemblyCode.WriteLine("// point A to the location pop from")
	assemblyCode.WriteLine("@SP")
	assemblyCode.WriteLine("A=M-1")

	assemblyCode.WriteLine("// the value from the stack")
	assemblyCode.WriteLine("D=M")

	assemblyCode.WriteLine("// write value to current local")
	assemblyCode.WriteLine("@LCL")
	assemblyCode.WriteLine("M=D")

	assemblyCode.WriteLine("// decrement SP")
	assemblyCode.WriteLine("@SP")
	assemblyCode.WriteLine("M=M-1")

	if command.Arg1 == "local" {
		assemblyCode.WriteLine("// reset LCL")
		assemblyCode.WriteLine("@LCL")
		assemblyCode.WriteLine("D=M")
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("M=D-A")
	}

	return string(assemblyCode) + "\n"
}
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
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("D=A") // Set D to Arg2
	case "local":
		assemblyCode.WriteLine("@LCL")
		assemblyCode.WriteLine("D=M")          // D is the base of local vars
		assemblyCode.WriteAValue(command.Arg2) // index from command
		assemblyCode.WriteLine("A=D+A")        // load the local address from the command
		assemblyCode.WriteLine("D=M")          // set D to the local value
	}

	// Add D to stack
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
func (ac *AssemblyCode) WriteAValue(value int) {
	ac.WriteLine("@" + strconv.Itoa(value))
}
