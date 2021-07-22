package extended

import "jvmgo/jvm/instructions/base"

// Branch always (wide index)
type GOTO_W struct {
	offset int
	Name   string
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) GetOperands() int {
	return self.offset
}

func (self *GOTO_W) GetName() string {
	return self.Name
}
