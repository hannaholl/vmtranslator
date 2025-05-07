
@256
D=A
@SP
M=D

@300
D=A
@LCL
M=D

@400
D=A
@ARG
M=D

@3000
D=A
@THIS
M=D

@3010
D=A
@THAT
M=D

// VM line: push constant 3030
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop pointer 0
// decrement stack and point to value to pop
@SP
AM=M-1
// the value from the stack
D=M
@THIS
M=D

// VM line: push constant 3040
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop pointer 1
// decrement stack and point to value to pop
@SP
AM=M-1
// the value from the stack
D=M
@THAT
M=D

// VM line: push constant 32
@32
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop this 2
// update segment to the new address
@2
D=A
@THIS
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@THIS
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment back to initial value
@THIS
D=M
@2
D=D-A
@THIS
M=D

// VM line: push constant 46
@46
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop that 6
// update segment to the new address
@6
D=A
@THAT
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@THAT
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment back to initial value
@THAT
D=M
@6
D=D-A
@THAT
M=D

// VM line: push pointer 0
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: push pointer 1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: add
@SP
AM=M-1
D=M
A=A-1
M=D+M

// VM line: push this 2
@THIS
D=M
@2
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: sub
@SP
AM=M-1
D=M
A=A-1
M=M-D

// VM line: push that 6
@THAT
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: add
@SP
AM=M-1
D=M
A=A-1
M=D+M

(END)
@END
0;JMP