
// VM line: push constant 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop static 8
// decrement stack and point to value to pop
@SP
AM=M-1
// the value from the stack
D=M
@StaticTest.8
M=D

// VM line: pop static 3
// decrement stack and point to value to pop
@SP
AM=M-1
// the value from the stack
D=M
@StaticTest.3
M=D

// VM line: pop static 1
// decrement stack and point to value to pop
@SP
AM=M-1
// the value from the stack
D=M
@StaticTest.1
M=D

// VM line: push static 3
@StaticTest.3
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: push static 1
@StaticTest.1
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

// VM line: push static 8
@StaticTest.8
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