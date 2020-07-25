package references

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 从当前类的运行时常量池找到一个类符号应用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析这个类符号引用，拿到类数据
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	// 创建对象
	ref := class.NewObject()
	// 把对象引用推入栈顶
	frame.OperandStack().PushRef(ref)
}
