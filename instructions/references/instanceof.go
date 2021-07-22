package references

import "jvmgo/jvm/instructions/base"

// Determine if object is of given type
type INSTANCE_OF struct{ base.Index16Instruction }
