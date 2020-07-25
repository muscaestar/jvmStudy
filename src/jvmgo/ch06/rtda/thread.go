package rtda

import "jvmgo/ch06/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack // JVM Stack
}

func NewThread() *Thread {
	return &Thread{
		// TODO 改进命令行工具，添加选项来指定栈的深度
		// 指定栈最大深度
		stack: newStack(1024),
	}
}

// getter
func (self *Thread) PC() int { return self.pc }

// setter
func (self *Thread) SetPC(pc int) { self.pc = pc }

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
