package base

type BytecodeReader struct {
	Code []byte // bytecodes
	Pc   int
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.Code = code
	self.Pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.Pc
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.Code[self.Pc]
	self.Pc++
	return i
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) SkipPadding() {
	for self.Pc%4 != 0 {
		self.ReadUint8()
	}
}

func (self *BytecodeReader) Finished() bool {
	return self.Pc >= len(self.Code)
}
