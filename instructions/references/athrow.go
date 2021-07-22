package references

import "jvmgo/jvm/instructions/base"

// Throw exception or error
type ATHROW struct{ base.NoOperandsInstruction }
