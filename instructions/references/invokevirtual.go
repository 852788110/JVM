package references

import (
	"jvmgo/jvm/instructions/base"
)

// Invoke instance method; dispatch based on class
type INVOKE_VIRTUAL struct{ base.Index16Instruction }
