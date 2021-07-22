package references

import (
	"jvmgo/jvm/instructions/base"
)
import "jvmgo/jvm/rtda"
import "jvmgo/jvm/rtda/heap"

// Set static field in class
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	// 如果没有开始初始化的话
	if !class.InitStarted() {
		// 如果没有初始化的话，那么此处设置开始初始化
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	} else if !class.InitFinished() {
		// 如果没有初始化完成的话
		// 如果初始化的线程是当前线程
		if class.GetThread() != frame.Thread().GetId() {
			class.GetMutex().Lock(frame.Thread().GetId())
		}
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
