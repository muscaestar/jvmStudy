package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// 0x10: bipush:  从操作数中获取一个byte型整数，扩展成int型，推入操作数栈顶
type BIPUSH struct{ val int8 } // Push byte
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// 0x11: sipush:  从操作数中获取一个short型整数，扩展成int型，推入操作数栈顶
type SIPUSH struct{ val int16 } // Push short
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
