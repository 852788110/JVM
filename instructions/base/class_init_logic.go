package base

import (
	"fmt"
	"jvmgo/jvm/rtda"
)
import "jvmgo/jvm/rtda/heap"

// jvms 5.5
func InitClass(thread *rtda.Thread, class *heap.Class) {
	if class.Name() == "java/lang/Integer" {
		fmt.Println("Nice")
	}
	class.GetMutex().Lock(thread)
	if !class.InitStarted() {
		scheduleClinit(thread, class)
		initSuperClass(thread, class)
		// class.StartInit()
	}
	// class.GetMutex().Unlock()
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
