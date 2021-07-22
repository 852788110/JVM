package references

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"
import "jvmgo/jvm/rtda/heap"

// Invoke a class (static) method
type INVOKE_STATIC struct{ base.Index16Instruction }

func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
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

	base.InvokeMethod(frame, resolvedMethod)
}
