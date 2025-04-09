
// VM line: push constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 200
@200
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop local 2
// update LCL to the new address
@2
D=A
@LCL
M=D+M
// get the location value to pop
@SP
A=M-1
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
@2
M=D-A

(END)
@END
0;JMP