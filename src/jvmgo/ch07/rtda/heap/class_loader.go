package heap

import (
	"fmt"
	"jvmgo/ch07/classfile"
	"jvmgo/ch07/classpath"
)

type Classloader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class // loaded classes
}

func NewClassloader(cp *classpath.Classpath) *Classloader {
	return &Classloader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// 把类数据加载到方法区
func (self *Classloader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class // 类已经加载
	}
	return self.loadNonArrayClass(name)
}

// 加载非数组类
func (self *Classloader) loadNonArrayClass(name string) *Class {
	// 读取class文件数据到内存
	data, entry := self.readClass(name)
	// 解析数据，生成虚拟机可以使用的类数据，并放入方法区
	class := self.defineClass(data)
	// 进行链接
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func link(class *Class) {
	verify(class)
	prepare(class) // 给类变量分配空间并给予初始值
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) // 计算实例字段的个数，并编号
	calcStaticFieldSlotIds(class)   // 计算静态字段的个数，并编号
	allocAndInitStaticVars(class)   // 给静态变量分配空间，然后赋予初始值
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVal(class, field)
		}
	}
}

// 从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVal(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// todo
			panic("todo")
		}
	}

}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func verify(class *Class) {
	// TODO 对类进行严格的验证，暂时忽略
}

func (self *Classloader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *Classloader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			// 递归调用加载方法加载直接接口
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		// 递归调用加载方法加载超类
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
