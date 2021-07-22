package references

import "jvmgo/jvm/instructions/base"

// Create new multidimensional array
type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
	Name       string
}

func (self *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) GetOperands() int {
	return -1
}

func (self *MULTI_ANEW_ARRAY) GetName() string {
	return self.Name
}
