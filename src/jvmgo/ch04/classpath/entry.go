package classpath

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找和加载class文件, className为相对路径
	// 返回值 读取的字节数据，最终定位的Entry，错误信息
	readClass(className string) ([]byte, Entry, error)

	String() string // 相当于Java的toString（）
}

func newEntry(path string) Entry {
	// 多个类路径
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
