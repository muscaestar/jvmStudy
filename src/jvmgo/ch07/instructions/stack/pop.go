package stack

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

// 用于弹出int、float等占用一个位置的变量
type POP struct{ base.NoOperandsInstruction }

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 用于弹出long、double等占用两个位置的变量
type POP2 struct{ base.NoOperandsInstruction }

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
