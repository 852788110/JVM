package references

import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"

// Enter monitor for object
type MONITOR_ENTER struct{ base.NoOperandsInstruction }

func (self *MONITOR_ENTER) Execute(frame *rtda.Frame) {
	// 得到monitor对象，对其加锁
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	monitor := ref.GetMonitor()
	monitor.Lock(frame.Thread().GetId())
}

// Exit monitor for object
type MONITOR_EXIT struct{ base.NoOperandsInstruction }

func (self *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	monitor := ref.GetMonitor()
	monitor.Unlock()
}
