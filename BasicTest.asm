
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

// VM line: push constant 10
@10
D=A
@SP
AM = M+1

A=M
M=D
@SP
M=M+1

// VM line: pop local 0
// update segment to the new address
@0
D=A
@LCL
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@LCL
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@LCL
D=M
@0
D=D-A
@LCL
M=D

// VM line: push constant 21
@21
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 22
@22
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop argument 2
// update segment to the new address
@2
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@2
D=D-A
@ARG
M=D

// VM line: pop argument 1
// update segment to the new address
@1
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@1
D=D-A
@ARG
M=D

// VM line: push constant 36
@36
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop this 6
// update segment to the new address
@6
D=A
@LCL
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@LCL
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@LCL
D=M
@6
D=D-A
@LCL
M=D

// VM line: push constant 42
@42
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push constant 45
@45
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop that 5
// update segment to the new address
@5
D=A
@LCL
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@LCL
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@LCL
D=M
@5
D=D-A
@LCL
M=D

// VM line: pop that 2
// update segment to the new address
@2
D=A
@LCL
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@LCL
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@LCL
D=M
@2
D=D-A
@LCL
M=D

// VM line: push constant 510
@510
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: pop temp 6
// update segment to the new address
@6
D=A
@LCL
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@LCL
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@LCL
D=M
@6
D=D-A
@LCL
M=D

// VM line: push local 0
@LCL
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// VM line: push that 5
@SP
A=M
M=D
@SP
M=M+1

// VM line: add

// update segment to the new address
@0
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@0
D=D-A
@ARG
M=D


// update segment to the new address
@1
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@1
D=D-A
@ARG
M=D


@15
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: push argument 1
@SP
A=M
M=D
@SP
M=M+1

// VM line: sub
// VM line: push this 6
@SP
A=M
M=D
@SP
M=M+1

// VM line: push this 6
@SP
A=M
M=D
@SP
M=M+1

// VM line: add

// update segment to the new address
@0
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@0
D=D-A
@ARG
M=D


// update segment to the new address
@1
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@1
D=D-A
@ARG
M=D


@15
D=A
@SP
A=M
M=D
@SP
M=M+1

// VM line: sub
// VM line: push temp 6
@SP
A=M
M=D
@SP
M=M+1

// VM line: add

// update segment to the new address
@0
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@0
D=D-A
@ARG
M=D


// update segment to the new address
@1
D=A
@ARG
M=D+M
// point A to the location pop from
@SP
A=M-1
// the value from the stack
D=M
// write value to current segment
@ARG
A=M
M=D
// decrement SP
@SP
M=M-1
// reset segment
@ARG
D=M
@1
D=D-A
@ARG
M=D


@15
D=A
@SP
A=M
M=D
@SP
M=M+1

(END)
@END
0;JMP