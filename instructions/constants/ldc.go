package constants

import "jvmgo/jvm/instructions/base"

// Push item from run-time constant pool
type LDC struct{ base.Index8Instruction }

// Push item from run-time constant pool (wide index)
type LDC_W struct{ base.Index16Instruction }

// Push long or double from run-time constant pool (wide index)
type LDC2_W struct{ base.Index16Instruction }
