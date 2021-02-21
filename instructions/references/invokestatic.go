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
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}