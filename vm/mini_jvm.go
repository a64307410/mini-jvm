package vm

import (
	"fmt"
	"github.com/wanghongfei/mini-jvm/vm/class"
	"os"
	"strings"
)

// VM定义
type MiniJvm struct {
	// 命令行参数
	CmdArgs []string

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

	ExecuteWithFrame(file *class.DefFile, methodName string, descriptor string, frame *MethodStackFrame) error
}

func NewMiniJvm(mainClass string, classPaths []string, cmdArgs... string) (*MiniJvm, error) {
	if "" == mainClass {
		return nil, fmt.Errorf("invalid main class '%s'", mainClass)
	}

	vmArgs := []string {os.Args[0]}

	if nil != cmdArgs {
		vmArgs = append(vmArgs, cmdArgs...)
	}

	vm := &MiniJvm{
		CmdArgs:  vmArgs,
		MethodArea: nil,
		MainClass:  strings.ReplaceAll(mainClass, ".", "/"),
		DebugPrintHistory: make([]interface{}, 0, 3),
	}

	// 方法区
	ma, err := NewMethodArea(vm, classPaths, nil)
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
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "print", "(I)V", PrintInt)
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "printInt", "(I)V", PrintInt)
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "printInt2", "(II)V", PrintInt2)
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "printChar", "(C)V", PrintChar)
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "printString", "(Ljava/lang/String;)V", PrintString)
	nativeMethodTable.RegisterMethod("cn.minijvm.io.Printer", "printBool", "(Z)V", PrintBoolean)

	nativeMethodTable.RegisterMethod("cn.minijvm.concurrency.MiniThread", "start", "(Ljava/lang/Runnable;)V", ExecuteInThread)
	nativeMethodTable.RegisterMethod("cn.minijvm.concurrency.MiniThread", "sleepCurrentThread", "(I)V", ThreadSleep)

	nativeMethodTable.RegisterMethod("java.lang.Object", "hashCode", "()I", ObjectHashCode)
	nativeMethodTable.RegisterMethod("java.lang.Object", "clone", "()Ljava/lang/Object;", ObjectClone)
	nativeMethodTable.RegisterMethod("java.lang.Object", "getClass", "()Ljava/lang/Class;", ObjectGetClass)

	nativeMethodTable.RegisterMethod("java.lang.Class", "getName0", "()Ljava/lang/String;", ClassGetName0)
	nativeMethodTable.RegisterMethod("java.lang.Class", "isInterface", "()Z", ClassIsInterface)
	nativeMethodTable.RegisterMethod("java.lang.Class", "isPrimitive", "()Z", ClassIsPrimitive)

	//public static native void arraycopy(Object src,  int  srcPos,
	//	Object dest, int destPos,
	//	int length);
	nativeMethodTable.RegisterMethod("java.lang.System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", SystemArrayCopy)

	return vm, nil
}

// 启动VM
func (m *MiniJvm) Start() error {
	return m.executeMain()
}

// 执行主类
func (m *MiniJvm) executeMain() error {
	mainClassDef, err := m.MethodArea.LoadClass(m.MainClass)
	if nil != err {
		return err
	}

	// 执行
	// log.Printf("main class info: %+v\n", mainClassDef)
	return m.ExecutionEngine.Execute(mainClassDef, "main")
}
