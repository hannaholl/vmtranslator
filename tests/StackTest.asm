
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

// VM line: push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 17
@17
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
D=M-D
@EQ_TRUE_10
D;JEQ
@SP
A=M-1
M=0
@EQ_END_10
0;JEQ
(EQ_TRUE_10)
@SP
A=M-1
M=1
(EQ_END_10)

// VM line: push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 16
@16
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
D=M-D
@EQ_TRUE_13
D;JEQ
@SP
A=M-1
M=0
@EQ_END_13
0;JEQ
(EQ_TRUE_13)
@SP
A=M-1
M=1
(EQ_END_13)

// VM line: push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 17
@17
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
D=M-D
@EQ_TRUE_16
D;JEQ
@SP
A=M-1
M=0
@EQ_END_16
0;JEQ
(EQ_TRUE_16)
@SP
A=M-1
M=1
(EQ_END_16)

// VM line: push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT_TRUE_19
D;JLT
@SP
A=M-1
M=0
@LT_END_19
0;JEQ
(LT_TRUE_19)
@SP
A=M-1
M=1
(LT_END_19)

// VM line: push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT_TRUE_22
D;JLT
@SP
A=M-1
M=0
@LT_END_22
0;JEQ
(LT_TRUE_22)
@SP
A=M-1
M=1
(LT_END_22)

// VM line: push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: lt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@LT_TRUE_25
D;JLT
@SP
A=M-1
M=0
@LT_END_25
0;JEQ
(LT_TRUE_25)
@SP
A=M-1
M=1
(LT_END_25)

// VM line: push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT_TRUE_28
D;JGT
@SP
A=M-1
M=0
@GT_END_28
0;JEQ
(GT_TRUE_28)
@SP
A=M-1
M=1
(GT_END_28)

// VM line: push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT_TRUE_31
D;JGT
@SP
A=M-1
M=0
@GT_END_31
0;JEQ
(GT_TRUE_31)
@SP
A=M-1
M=1
(GT_END_31)

// VM line: push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: gt
@SP
AM=M-1
D=M
A=A-1
D=M-D
@GT_TRUE_34
D;JGT
@SP
A=M-1
M=0
@GT_END_34
0;JEQ
(GT_TRUE_34)
@SP
A=M-1
M=1
(GT_END_34)

// VM line: push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 53
@53
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

// VM line: push constant 112
@112
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
M=M-D

// VM line: neg
@SP
A=M-1
M=-M

// VM line: and
@SP
AM=M-1
D=M
A=A-1
M=D&M

// VM line: push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: or
@SP
AM=M-1
D=M
A=A-1
M=D|M

// VM line: not
@SP
A=M-1
M=!M

(END)
@END
0;JMP