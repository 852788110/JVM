package stack

import "jvmgo/jvm/instructions/base"

// Duplicate the top operand stack value
type DUP struct{ base.NoOperandsInstruction }

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperandsInstruction }

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperandsInstruction }

// Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperandsInstruction }

// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperandsInstruction }

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperandsInstruction }
