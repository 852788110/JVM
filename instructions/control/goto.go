package control

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"

// Branch always
type GOTO struct{ base.BranchInstruction }

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
