package main

import (
	"jvmgo/jvm/native"
	"jvmgo/jvm/rtda"
)

func init() {
	native.Register("NewThread", "start0", "()V",start0)
}


// 参数: 一个thread对象的引用
// 创建一个jvm虚拟机栈
// 将thread对象的run方法入栈
func start0(frame *rtda.Frame) {
	// 得到其class对象
	currentClass := frame.Method().Class()

	subThread := rtda.NewThread()

	newframe := rtda.NewFrame(subThread, currentClass.GetMethod("run", "()V", false))

	subThread.PushFrame(newframe)

	go interpret(subThread, false)

}
