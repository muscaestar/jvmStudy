package base

import "jvmgo/ch07/rtda"

// 把指令抽象成接口
type Instruction interface {
	// 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// 执行指令逻辑
	Execute(frame *rtda.Frame)
}

// 没有操作数的指令
type NoOperandsInstruction struct{}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (self *NoOperandsInstruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}

// 跳转指令
type BranchInstruction struct {
	Offset int // 跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *BranchInstruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}

// 存储和加载类指令
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (self *Index8Instruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}

// 存储和加载类指令
type Index16Instruction struct {
	Index uint // 常量池索引
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

func (self *Index16Instruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}
