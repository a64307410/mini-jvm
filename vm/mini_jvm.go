package vm

import (
	"fmt"
	"github.com/wanghongfei/mini-jvm/vm/class"
	"strings"
)

// VM定义
type MiniJvm struct {
	// 方法区
	MethodArea *MethodArea

	// MainClass全限定性名
	MainClass string

	// 执行引擎
	ExecutionEngine ExecutionEngine

	// 本地方法表
	NativeMethodTable *NativeMethodTable

	// 保存调用print的历史记录, 单元测试用
	DebugPrintHistory []interface{}
}

type ExecutionEngine interface {
	Execute(file *class.DefFile, methodName string) error

	ExecuteWithDescriptor(file *class.DefFile, methodName string, descriptor string) error
}

func NewMiniJvm(mainClass string, classPaths[] string) (*MiniJvm, error) {
	if "" == mainClass {
		return nil, fmt.Errorf("invalid main class '%s'", mainClass)
	}


	vm := &MiniJvm{
		MethodArea: nil,
		MainClass:  strings.ReplaceAll(mainClass, ".", "/"),
		DebugPrintHistory: make([]interface{}, 0, 3),
	}

	// 方法区
	ma, err := NewMethodArea(vm, classPaths)
	if nil != err {
		return nil, fmt.Errorf("unabled to create method area: %w", err)
	}
	vm.MethodArea = ma

	// 执行引擎
	vm.ExecutionEngine = NewInterpretedExecutionEngine(vm)

	// 本地方法表
	nativeMethodTable := NewNativeMethodTable()
	vm.NativeMethodTable = nativeMethodTable
	// 注册本地方法
	nativeMethodTable.RegisterMethod("print", "(I)V", PrintInt)
	nativeMethodTable.RegisterMethod("printInt", "(I)V", PrintInt)
	nativeMethodTable.RegisterMethod("printInt2", "(II)V", PrintInt2)
	nativeMethodTable.RegisterMethod("printChar", "(C)V", PrintChar)

	return vm, nil
}

// 启动VM
func (m *MiniJvm) Start() error {
	return m.executeMain()
}

// 执行主类
func (m *MiniJvm) executeMain() error {
	mainClassDef, err := m.findDefClass(m.MainClass)
	if nil != err {
		return err
	}

	// 执行
	// log.Printf("main class info: %+v\n", mainClassDef)
	return m.ExecutionEngine.Execute(mainClassDef, "main")
}

func (m *MiniJvm) findDefClass(className string) (*class.DefFile, error) {
	// 从已加载的类中查找
	def, ok := m.MethodArea.ClassMap[className]
	if ok {
		return def, nil
	}

	// 不存在, 触发加载
	def, err := m.MethodArea.LoadClass(className)
	if nil != err {
		return nil, fmt.Errorf("unabled to load class '%s': %w", className, err)
	}

	return def, nil
}
