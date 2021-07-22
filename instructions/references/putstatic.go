package references

import (
	"jvmgo/jvm/instructions/base"
)

// Set static field in class
type PUT_STATIC struct{ base.Index16Instruction }
