package references

import "jvmgo/jvm/instructions/base"

// Enter monitor for object
type MONITOR_ENTER struct{ base.NoOperandsInstruction }

// Exit monitor for object
type MONITOR_EXIT struct{ base.NoOperandsInstruction }
