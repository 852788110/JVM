package references

import "jvmgo/jvm/instructions/base"

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }
