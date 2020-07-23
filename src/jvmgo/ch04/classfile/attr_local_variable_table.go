package classfile

type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc       uint16
	localVariable uint16
}

func (self *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	LocalVariableTableLength := reader.readUint16()
	self.LocalVariableTable = make([]*LocalVariableTableEntry, LocalVariableTableLength)
	for i := range self.LocalVariableTable {
		self.LocalVariableTable[i] = &LocalVariableTableEntry{
			startPc:       reader.readUint16(),
			localVariable: reader.readUint16(),
		}
	}
}
