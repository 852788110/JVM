package math

import "jvmgo/jvm/instructions/base"

// Add double
type DADD struct{ base.NoOperandsInstruction }

// Add float
type FADD struct{ base.NoOperandsInstruction }

// Add int
type IADD struct{ base.NoOperandsInstruction }

// Add long
type LADD struct{ base.NoOperandsInstruction }
