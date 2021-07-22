package base

type Instruction interface {
	GetOperands() int
	GetName() string
	FetchOperands(reader *BytecodeReader)
}

type NoOperandsInstruction struct {
	// empty
	Name string
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (self *NoOperandsInstruction) GetOperands() int {
	// nothing to do
	return -1
}

func (self *NoOperandsInstruction) GetName() string {
	// nothing to do
	return self.Name
}

type BranchInstruction struct {
	Offset int
	Name   string
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *BranchInstruction) GetOperands() int {
	return self.Offset
}

func (self *BranchInstruction) GetName() string {
	// nothing to do
	return self.Name
}

type Index8Instruction struct {
	Index uint
	Name  string
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (self *Index8Instruction) GetOperands() int {
	return int(self.Index)
}

func (self *Index8Instruction) GetName() string {
	// nothing to do
	return self.Name
}

type Index16Instruction struct {
	Index uint
	Name  string
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

func (self *Index16Instruction) GetOperands() int {
	return int(self.Index)
}

func (self *Index16Instruction) GetName() string {
	// nothing to do
	return self.Name
}
