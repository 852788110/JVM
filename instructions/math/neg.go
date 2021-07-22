package math

import "jvmgo/jvm/instructions/base"

// Negate double
type DNEG struct{ base.NoOperandsInstruction }

// Negate float
type FNEG struct{ base.NoOperandsInstruction }

// Negate int
type INEG struct{ base.NoOperandsInstruction }

// Negate long
type LNEG struct{ base.NoOperandsInstruction }
