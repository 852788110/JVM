package stack

import "jvmgo/jvm/instructions/base"

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }
