package base

import (
	"jvmgo/jvm/rtda"
)
import "jvmgo/jvm/rtda/heap"

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.GetMutex().Lock(thread.GetId())
	if !class.InitStarted() {
		class.StartInit()
		scheduleClinit(thread, class)
		initSuperClass(thread, class)
	}
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil && clinit.Class() == class {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}

func InitClassMutil(thread *rtda.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}
