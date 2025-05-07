
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

// VM line: push constant 111
@111
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

// VM line: eq
@SP
AM=M-1
D=M
A=A-1
D=D-M
@EQ_TRUE_9
D;JEQ
@SP
A=M-1
M=0
@EQ_END_9
0;JEQ
(EQ_TRUE_9)
@SP
A=M-1
M=1
(EQ_END_9)

(END)
@END
0;JMP