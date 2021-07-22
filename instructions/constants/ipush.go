package constants

import "jvmgo/jvm/instructions/base"

// Push byte
type BIPUSH struct {
	val  int8
	Name string
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) GetOperands() int {
	return int(self.val)
}

func (self *BIPUSH) GetName() string {
	return self.Name
}

// Push short
type SIPUSH struct {
	val  int16
	Name string
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) GetOperands() int {
	return int(self.val)
}

func (self *SIPUSH) GetName() string {
	return self.Name
}
