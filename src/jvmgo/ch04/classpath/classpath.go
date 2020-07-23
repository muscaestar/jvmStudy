package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 启动类路径 -Xjre
	extClasspath  Entry // 扩展类路径 -Xjre
	userClasspath Entry // 用户类路径 -cp/classpath
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*") // <jre>/lib/*
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*") // <jre>/lib/ext/*
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	// 优先使用-Xjre选项作为jre目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 在当前目录下寻找jre
	if exists("./jre") {
		return "./jre"
	}
	// 使用JAVA_HOME环境变量获取jdk路径
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
