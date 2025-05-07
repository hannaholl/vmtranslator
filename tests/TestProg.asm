
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

// VM line: push constant 24
@24
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: not
@SP
AM=M-1
M=!M

(END)
@END
0;JMP