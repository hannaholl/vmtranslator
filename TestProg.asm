
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

// VM line: push constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 111
@111
D=A
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
M=D-M

// VM line: push constant 444
@444
D=A
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