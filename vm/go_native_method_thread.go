package vm

import (
	"fmt"
	"github.com/wanghongfei/mini-jvm/vm/class"
	"time"
)

const (
	THREAD_STATUS_CREATED = 0
	THREAD_STATUS_RUNNING = 0
	THREAD_STATUS_FINISHED = 0
)

// java线程对应go里的表示
type MiniThread struct {
	Jvm *MiniJvm
	JavaObjRef *class.Reference

	// 线程状态
	// 0: created
	// 1: running
	// 2: finished
	Status int
}

func (t *MiniThread) Start() {
	// 创建栈帧
	// 把objRef压进去
	opStack := NewOpStack(1)
	opStack.Push(t.JavaObjRef)
	frame := &MethodStackFrame{
		localVariablesTable: nil,
		opStack:             opStack,
		pc:                  0,
	}

	go func() {
		t.Status = THREAD_STATUS_RUNNING

		// 防止进程崩溃
		defer func() {
			r := recover()
			if nil != r {
				fmt.Printf("goroutine recovered: %v\n", r)
			}
		}()

		defer func() {
			t.Status = THREAD_STATUS_FINISHED
		}()

		err := t.Jvm.ExecutionEngine.ExecuteWithFrame(t.JavaObjRef.Object.DefFile, "run", "()V", frame)
		if nil != err {
			fmt.Printf("failed to execute native function 'ExecuteInThread': %v\n", err)
		}
	}()
}

// 当前协程sleep指定秒数
func ThreadSleep(args ...interface{}) interface{} {
	seconds := args[1].(int)
	time.Sleep(time.Duration(seconds) * time.Second)

	return true
}

// 在新的协程中执行字节码
func ExecuteInThread(args ...interface{}) interface{} {
	// 第一个参数为jvm指针
	jvm := args[0].(*MiniJvm)
	// 第二个参数是实现了Runnalbe接口的对象引用
	objRef := args[1].(*class.Reference)

	miniThread := &MiniThread{
		Jvm:        jvm,
		JavaObjRef: objRef,
		Status: THREAD_STATUS_CREATED,
	}
	miniThread.Start()

	return true
}
