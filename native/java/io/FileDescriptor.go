package io

import "jvmgo/jvm/native"
import "jvmgo/jvm/rtda"

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
}

// private static native long set(int d);
// (I)J
func set(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
