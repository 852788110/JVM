package references

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"
import "jvmgo/jvm/rtda/heap"

// Get static field from class
type GET_STATIC struct{ base.Index16Instruction }

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
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

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
