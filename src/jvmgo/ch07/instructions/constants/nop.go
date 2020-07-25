package constants

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

// 0x00: nop: do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
