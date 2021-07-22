package stack

import "jvmgo/jvm/instructions/base"

// Pop the top operand stack value
type POP struct{ base.NoOperandsInstruction }

// Pop the top one or two operand stack values
type POP2 struct{ base.NoOperandsInstruction }
