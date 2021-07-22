package main

import (
	"jvmgo/jvm/native"
	"jvmgo/jvm/rtda"
	"sync"
)

// common/Print.print0(Ljava/lang/String;)V
func init() {
	native.Register("common/Thread", "start0", "()V", start0)
}

var ThreadIds int = 1

// 同步创建线程的互斥量
var threadMutex sync.Mutex

// 参数: 一个thread对象的引用
// 创建一个jvm虚拟机栈
// 将thread对象的run方法入栈
func start0(frame *rtda.Frame) {
	// 得到其class对象
	//currentClass := frame.Method().Class()

	subThread := rtda.NewThread()

	// 为该子线程分配一个ThreadId，这个地方得加锁
	threadMutex.Lock()
	subThread.SetId(ThreadIds)
	ThreadIds++
	threadMutex.Unlock()

	// 从前一个栈帧的操作数栈上取得　对象引用
	ref := frame.LocalVars().GetThis()
	newframe := rtda.NewFrame(subThread, ref.Class().GetMethod("run", "()V", false))

	// 将对象引用放到新创键的frame的0号位置
	newframe.LocalVars().SetRef(0, ref)

	subThread.PushFrame(newframe)

	// 防止main进程退出了之后，所有的groutine都结束了
	Wg.Add(1)
	go interpret(subThread, false)
}
