package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/instructions"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	// 获取方法的Code属性
	codeAttr := methodInfo.CodeAttribute()
	// 获取执行方法所需的局部变量表空间
	maxLocals := codeAttr.MaxLocals()
	// 获取执行方法所需的操作数栈空间
	maxStack := codeAttr.MaxStack()
	// 获取方法的字节码
	bytecode := codeAttr.Code()
	// 创建一个Thread
	thread := rtda.NewThread()
	// 给Thread创建第一个栈帧
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	// 循环执行：计算pc、解码指令、执行指令
	loop(thread, bytecode)
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
