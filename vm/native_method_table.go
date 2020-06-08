package vm

import (
	"strings"
)

type NativeFunction func(args ...interface{}) interface{}

type NativeMethodInfo struct {
	// 方法名
	Name string

	// 描述符;
	// String getRealnameByIdAndNickname(int id,String name) 的描述符为 (ILjava/lang/String;)Ljava/lang/String;
	Descriptor string

	// 对应的go函数
	EntryFunc NativeFunction
}

// 解析方法描述符, 返回方法参数的数量
// 注意, 不支持方法描述符中含有对象类型
func (info *NativeMethodInfo) ParseArgCount() int {
	argAndRetDesciptor := strings.Split(info.Descriptor, ")")
	argDescriptor := argAndRetDesciptor[0][1:]

	return len(argDescriptor)

	// return strings.Count(argDescriptor, ",") + 1
}

// 本地方法表
type NativeMethodTable struct {
	MethodInfoMap map[string]*NativeMethodInfo
}

func NewNativeMethodTable() *NativeMethodTable {
	return &NativeMethodTable{MethodInfoMap: map[string]*NativeMethodInfo{}}
}

// 注册本地方法
// methodName: 方法名
// descriptor: 方法在JVM中的描述符
func (t *NativeMethodTable) RegisterMethod(methodName string, descriptor string, goFunc NativeFunction) {
	key := t.genKey(methodName, descriptor)
	t.MethodInfoMap[key] = &NativeMethodInfo{
		Name:       methodName,
		Descriptor: descriptor,
		EntryFunc:  goFunc,
	}
}

// 查本地方法表, 找出目标go函数
func (t *NativeMethodTable) FindMethod(name string, descriptor string) (NativeFunction, int) {
	f, ok := t.MethodInfoMap[t.genKey(name, descriptor)]
	if !ok {
		return nil, -1
	}

	return f.EntryFunc, f.ParseArgCount()
}


func (t *NativeMethodTable) genKey(name string, descriptor string) string {
	return name + "=>" + descriptor
}
