package codewriter

import (
	"nandtotetris/vmtranslator/parser"
	"strconv"
)

type AssemblyCode string

func WriteArithmetic(command parser.Command, lineNumber int) string {
	var assemblyCode AssemblyCode = ""

	assemblyCode.WriteLine("@SP")

	switch command.Arg1 {
	case "add":
		assemblyCode.WriteLine("AM=M-1") // point to first operand an update SP
		assemblyCode.WriteLine("D=M")    // D is the first operand
		assemblyCode.WriteLine("A=A-1")  // Move to previous location in stack
		assemblyCode.WriteLine("M=D+M")  // Write computed value to stack
	case "sub":
		assemblyCode.WriteLine("AM=M-1") // point to first operand an update SP
		assemblyCode.WriteLine("D=M")    // D is the first operand
		assemblyCode.WriteLine("A=A-1")  // Move to previous location in stack
		assemblyCode.WriteLine("M=M-D")  // Write computed value to stack
	case "neg":
		assemblyCode.WriteLine("A=M-1") // point to operand, don't change SP
		assemblyCode.WriteLine("M=-M")  // negate the value
	case "eq":
		assemblyCode.WriteLine("AM=M-1") // point to first operand an update SP
		assemblyCode.WriteLine("D=M")    // D is the first operand
		assemblyCode.WriteLine("A=A-1")  // Move to previous location in stack
		assemblyCode.WriteLine("D=D-M")

		lineId := strconv.Itoa(lineNumber)
		trueLabel := "EQ_TRUE_" + lineId
		endLabel := "EQ_END_" + lineId

		// if D==0, jump to equal true
		assemblyCode.WriteLine("@" + trueLabel)
		assemblyCode.WriteLine("D;JEQ")

		// // if not jumping, write false
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("A=M-1")
		assemblyCode.WriteLine("M=0")

		// go to end
		assemblyCode.WriteLine("@" + endLabel)
		assemblyCode.WriteLine("0;JEQ")

		assemblyCode.WriteLine("(" + trueLabel + ")")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("A=M-1")
		assemblyCode.WriteLine("M=1")

		assemblyCode.WriteLine("(" + endLabel + ")")
	}

	return string(assemblyCode) + "\n"
}

func WritePop(command parser.Command, staticFileName string) string {
	var assemblyCode AssemblyCode = ""

	switch command.Arg1 {
	case "pointer":
		assemblyCode.WriteLine("// decrement stack and point to value to pop")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("AM=M-1")
		assemblyCode.WriteLine("// the value from the stack")
		assemblyCode.WriteLine("D=M")

		if command.Arg2 == 0 {
			assemblyCode.WriteLine("@THIS")
		} else {
			assemblyCode.WriteLine("@THAT")
		}
		assemblyCode.WriteLine("M=D")
	case "static":
		varName := staticFileName + "." + strconv.Itoa(command.Arg2)

		assemblyCode.WriteLine("// decrement stack and point to value to pop")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("AM=M-1")

		assemblyCode.WriteLine("// the value from the stack")
		assemblyCode.WriteLine("D=M")

		assemblyCode.WriteLine("@" + varName)
		assemblyCode.WriteLine("M=D")
	case "temp":
		// temp starts at 5
		tempLocation := strconv.Itoa(5 + command.Arg2)

		assemblyCode.WriteLine("// decrement stack and point to value to pop")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("AM=M-1")

		assemblyCode.WriteLine("// the value from the stack")
		assemblyCode.WriteLine("D=M")

		assemblyCode.WriteLine("@" + tempLocation)
		assemblyCode.WriteLine("M=D")
	default:
		segmentName := segmentMap[command.Arg1]

		assemblyCode.WriteLine("// update segment to the new address")
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("D=A")
		assemblyCode.WriteLine("@" + segmentName)
		assemblyCode.WriteLine("M=D+M")

		assemblyCode.WriteLine("// point A to the location pop from")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("A=M-1")

		assemblyCode.WriteLine("// the value from the stack")
		assemblyCode.WriteLine("D=M")

		assemblyCode.WriteLine("// write value to current segment")
		assemblyCode.WriteLine("@" + segmentName)
		assemblyCode.WriteLine("A=M")
		assemblyCode.WriteLine("M=D")

		assemblyCode.WriteLine("// decrement SP")
		assemblyCode.WriteLine("@SP")
		assemblyCode.WriteLine("M=M-1")

		assemblyCode.WriteLine("// reset segment back to initial value")
		assemblyCode.WriteLine("@" + segmentName)
		assemblyCode.WriteLine("D=M")
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("D=D-A")
		assemblyCode.WriteLine("@" + segmentName)
		assemblyCode.WriteLine("M=D")
	}

	return string(assemblyCode) + "\n"
}
func WritePush(command parser.Command, staticFileName string) string {

	var assemblyCode AssemblyCode = ""

	// Could write directly to the file instead of a variable?
	switch command.Arg1 {
	case "pointer":
		if command.Arg2 == 0 {
			assemblyCode.WriteLine("@THIS")
		} else {
			assemblyCode.WriteLine("@THAT")
		}
		assemblyCode.WriteLine("D=M")
	case "static":
		varName := staticFileName + "." + strconv.Itoa(command.Arg2)

		assemblyCode.WriteLine("@" + varName)
		assemblyCode.WriteLine("D=M") // D is the value to push
	case "constant":
		assemblyCode.WriteAValue(command.Arg2)
		assemblyCode.WriteLine("D=A") // Set D to Arg2
	case "temp":
		// temp starts at 5 but does not have a symbol
		assemblyCode.WriteLine("@" + strconv.Itoa(5+command.Arg2))
		assemblyCode.WriteLine("D=M") // D is the value to push
	default:
		assemblyCode.WriteLine("@" + segmentMap[command.Arg1])
		assemblyCode.WriteLine("D=M")          // D is the base of argument vars
		assemblyCode.WriteAValue(command.Arg2) // index from command
		assemblyCode.WriteLine("A=D+A")        // load the local address from the command
		assemblyCode.WriteLine("D=M")          // set D to the local value
	}

	// Add D to stack
	assemblyCode.WriteLine("@SP")

	// alternative approach. it's less code but for now I find it clearer
	// to do it in steps and increment the stack on its own at the end
	// assemblyCode.WriteLine("AM=M+1") // increment stack pointer
	// assemblyCode.WriteLine("A=A-1")  // point to where to push new value
	// assemblyCode.WriteLine("M=D")    // Set the current memory location (A) to D

	assemblyCode.WriteLine("A=M") // A points to the top of the stack
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

func WriteVariable(variableName string, value int) string {
	var assemblyCode AssemblyCode = ""
	assemblyCode.WriteAValue(value)
	assemblyCode.WriteLine("D=A")
	assemblyCode.WriteLine("@" + variableName)
	assemblyCode.WriteLine("M=D")

	return string(assemblyCode) + "\n"
}

var segmentMap = map[string]string{
	"argument": "ARG",
	"local":    "LCL",
	"this":     "THIS",
	"that":     "THAT",
}
