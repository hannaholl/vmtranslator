// VM line: push local 10
@LCL
D=M
@10
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
(END)
@END
0;JMP