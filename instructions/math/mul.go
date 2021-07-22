package math

import "jvmgo/jvm/instructions/base"

// Multiply double
type DMUL struct{ base.NoOperandsInstruction }

// Multiply float
type FMUL struct{ base.NoOperandsInstruction }

// Multiply int
type IMUL struct{ base.NoOperandsInstruction }

// Multiply long
type LMUL struct{ base.NoOperandsInstruction }
