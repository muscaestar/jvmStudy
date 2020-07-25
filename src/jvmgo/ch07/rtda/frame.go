package rtda

import "jvmgo/ch07/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 线程
	method       *heap.Method
	nextPC       int
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

// setters
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
