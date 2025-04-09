// pop local 4

// update LCL to the new address
@4
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
@4
M=D-A
