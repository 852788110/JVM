package model

type Class struct {
	//AccessFlags    uint16
	Name           string // thisClassName
	SuperClassName string
	InterfaceNames []string
	//ConstantPool   *heap.ConstantPool
	Fields  []*Field
	Methods []*Method
}

type Field struct {
	Name       string
	Descriptor string
}

type Method struct {
	Name       string
	Descriptor string
	Code       *Code
}

type Code struct {
	Instructions []string
}
