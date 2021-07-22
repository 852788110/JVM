package math

import "jvmgo/jvm/instructions/base"

// Shift left int
type ISHL struct{ base.NoOperandsInstruction }

// Arithmetic shift right int
type ISHR struct{ base.NoOperandsInstruction }

// Logical shift right int
type IUSHR struct{ base.NoOperandsInstruction }

// Shift left long
type LSHL struct{ base.NoOperandsInstruction }

// Arithmetic shift right long
type LSHR struct{ base.NoOperandsInstruction }

// Logical shift right long
type LUSHR struct{ base.NoOperandsInstruction }
