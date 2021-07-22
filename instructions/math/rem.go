package math

import "jvmgo/jvm/instructions/base"

// Remainder double
type DREM struct{ base.NoOperandsInstruction }

// Remainder float
type FREM struct{ base.NoOperandsInstruction }

// Remainder int
type IREM struct{ base.NoOperandsInstruction }

// Remainder long
type LREM struct{ base.NoOperandsInstruction }
