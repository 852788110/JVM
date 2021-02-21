package main

import "fmt"
import "jvmgo/jvm/instructions"
import "jvmgo/jvm/instructions/base"
import "jvmgo/jvm/rtda"

// 解释器
func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

/*
	do{
		计算pc寄存器以及从pc寄存器的位置取出操作码
		如果存在操作码，取出操作码
		执行操作码所定义的操作
	}
*/
func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		// 得到当前的操作数栈帧
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 读取栈帧中code属性
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		//　得到下一条指令
		inst := instructions.NewInstruction(opcode)
		// 取出操作
		inst.FetchOperands(reader)
		//　计算下一条指令的pc值
		frame.SetNextPC(reader.PC())

		// Debug模式
		if logInst {
			logInstruction(frame, inst)
		}

		// 执行指令
		inst.Execute(frame)
		// 如果jvm为空，退出执行
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
