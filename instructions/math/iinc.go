package math

import "jvmgo/jvm/instructions/base"

// Increment local variable by constant
type IINC struct {
	Index uint
	Const int32
	Name  string
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) GetOperands() int {
	return -1
}

func (self *IINC) GetName() string {
	return self.Name
}
