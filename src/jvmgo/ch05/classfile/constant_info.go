package classfile

// tag常量值定义，共14种
const (
	// 字面量：数字常量 + 字符串常量
	// 字符串常量
	CONSTANT_Utf8 = 1
	//	数字常量
	CONSTANT_Integer = 3
	CONSTANT_Float   = 4
	CONSTANT_Long    = 5
	CONSTANT_Double  = 6

	// 符号引用：类和接口名称、字段和方法的名称及描述符
	CONSTANT_Class              = 7  // ->utf8
	CONSTANT_String             = 8  // ->utf8
	CONSTANT_Fieldref           = 9  // -> class,nameAndType
	CONSTANT_Methodref          = 10 // -> class,nameAndType
	CONSTANT_InterfaceMethodref = 11 // -> class,nameAndType
	CONSTANT_NameAndType        = 12 // ->utf8,utf8

	// 支持invoke-dynamic指令，暂不做了解
	CONSTANT_MethodHandle  = 15
	CONSTANT_MethodType    = 16
	CONSTANT_InvokeDynamic = 18
)
