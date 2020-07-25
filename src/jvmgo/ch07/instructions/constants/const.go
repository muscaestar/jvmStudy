package constants

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type ACONST_NULL struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }

// 0x01: aconst_null: 把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 0x02: iconst_m1:  把int型 -1 推入操作数栈
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 0x03: iconst_0:  把int型 0 推入操作数栈
func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// 0x04: iconst_1:  把int型 1 推入操作数栈
func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

// 0x05: iconst_2:  把int型 2 推入操作数栈
func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

// 0x06: iconst_3:  把int型 3 推入操作数栈
func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

// 0x07: iconst_4:  把int型 4 推入操作数栈
func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

// 0x08: iconst_5:  把int型 5 推入操作数栈
func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// 0x09: lconst_0:  把long型 0 推入操作数栈顶
func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(int64(0))
}

// 0x0A: lconst_1:  把long型 1 推入操作数栈顶
func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(int64(1))
}

// 0x0B: fconst_0:  把float型 0 推入操作数栈顶
func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(float32(0.0))
}

// 0x0C: fconst_1:  把float型 1 推入操作数栈顶
func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(float32(1.0))
}

// 0x0D: fconst_2:  把float型 2 推入操作数栈顶
func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(float32(2.0))
}

// 0x0E: dconst_0： 把double型 0 推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 0x0F: dconst_1： 把double型 1 推入操作数栈顶
func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// 0x12 ldc

// 0x13 ldc_w

// 0x14 ldc2_w
