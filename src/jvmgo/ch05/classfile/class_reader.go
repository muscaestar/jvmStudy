package classfile

import "encoding/binary"

// 是[]byte类型的包装
type ClassReader struct {
	data []byte
}

// 读取class文件中类型为u1的数据
func (self *ClassReader) readUint8() uint8 {
	// 读取字节流中的第一个字节
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 读取class文件中类型为u2的数据
func (self *ClassReader) readUint16() uint16 {
	// 读取大端方式存储的数据，读字节流的前2个字节
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读取class文件中类型为u4的数据
func (self *ClassReader) readUint32() uint32 {
	// 读取大端方式存储的数据，读字节流的前4个字节
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	// 读取大端方式存储的数据，读字节流的前8个字节
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取class文件中类型为u2的数据表
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() // 以开头的2字节作为表的大小
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// 读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	// 读取字节流的前n个字节
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
