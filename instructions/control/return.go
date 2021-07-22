package control

import (
	"jvmgo/jvm/instructions/base"
)

// Return void from method
type RETURN struct{ base.NoOperandsInstruction }

// Return reference from method
type ARETURN struct{ base.NoOperandsInstruction }

// Return double from method
type DRETURN struct{ base.NoOperandsInstruction }

// Return float from method
type FRETURN struct{ base.NoOperandsInstruction }

// Return int from method
type IRETURN struct{ base.NoOperandsInstruction }

// Return double from method
type LRETURN struct{ base.NoOperandsInstruction }
