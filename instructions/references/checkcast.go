package references

import "jvmgo/jvm/instructions/base"

// Check whether object is of given type
type CHECK_CAST struct{ base.Index16Instruction }
