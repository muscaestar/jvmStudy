package main

import (
	"fmt"
	"jvmgo/ch06/instructions"
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
	"jvmgo/ch06/rtda/heap"
)

func interpret(method *heap.Method) {
	// 创建一个Thread
	thread := rtda.NewThread()
	// 给Thread创建第一个栈帧
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	// 循环执行：计算pc、解码指令、执行指令
	loop(thread, method.Code())
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}

}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
