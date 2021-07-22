package lang

import (
	"fmt"
	"jvmgo/jvm/native"
)
import "jvmgo/jvm/rtda"
import "jvmgo/jvm/rtda/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
	native.Register("common/Fmt", "print0", "(Ljava/lang/String;)V", print0)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}

func print0(frame *rtda.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	fmt.Println(name)
}
