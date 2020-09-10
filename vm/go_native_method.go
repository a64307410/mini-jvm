package vm

import (
	"fmt"
	"github.com/wanghongfei/mini-jvm/vm/class"
	"sync"
)

func PrintInt(args ...interface{}) interface{} {
	fmt.Println(args[2])

	return nil
}

func PrintInt2(args ...interface{}) interface{} {
	fmt.Println(args[2])
	fmt.Println(args[3])

	return nil
}

func PrintChar(args ...interface{}) interface{} {
	fmt.Printf("%c\n", args[2])

	return nil
}

func PrintString(args ...interface{}) interface{} {
	strRef := args[2].(*class.Reference)
	field := strRef.Object.ObjectFields["value"]
	strVal := field.FieldValue.([]rune)

	fmt.Printf("%v\n", string(strVal))

	return nil
}

// Object.hashcode()方法实现
// return: int
func ObjectHashCode(args ...interface{}) interface{} {
	ref := args[1].(*class.Reference)
	return ref.Object.HashCode
}

// Object.clone()方法实现
func ObjectClone(args ...interface{}) interface{} {
	// 要克隆的对象的引用
	targetRef := args[1].(*class.Reference)

	targetObj := &class.Object{
		DefFile:      targetRef.Object.DefFile,
		HashCode:     targetRef.Object.HashCode + 1,
		ObjectFields: targetRef.Object.ObjectFields,
	}

	newRef := &class.Reference{
		RefType: targetRef.RefType,
		Object:  targetObj,
		Array:   nil,
		Monitor: sync.Mutex{},
	}

	return newRef
}