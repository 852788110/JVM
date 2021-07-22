package math

import "jvmgo/jvm/instructions/base"

// Subtract double
type DSUB struct{ base.NoOperandsInstruction }

// Subtract float
type FSUB struct{ base.NoOperandsInstruction }

// Subtract int
type ISUB struct{ base.NoOperandsInstruction }

// Subtract long
type LSUB struct{ base.NoOperandsInstruction }
