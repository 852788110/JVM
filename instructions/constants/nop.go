package constants

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
