package references

import "jvmgo/jvm/instructions/base"

// Invoke interface method
type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
	Name string
}

func (self *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	self.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // must be 0
}

func (self *INVOKE_INTERFACE) GetOperands() int {
	return int(self.index)
}

func (self *INVOKE_INTERFACE) GetName() string {
	return self.Name
}
