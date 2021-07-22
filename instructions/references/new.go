package references

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"
import "jvmgo/jvm/rtda/heap"

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

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

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
