package vm

import (
	"testing"
)

func TestHelloNative(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.ForLoopPrintTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}
}

func TestHelloClass(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.NewSimpleObjectTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}

}

func TestHelloMethod(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.MethodReloadTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}

}

func TestClassExtend(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.ClassExtendTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}
}

func TestRecursion(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.RecursionTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}
}

func TestIfTest(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.IfTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}
}

func TestIntArray(t *testing.T) {
	miniJvm, err := NewMiniJvm("com.fh.ArrayTest", []string{"../testclass/"})
	if nil != err {
		t.Fatal(err)
	}

	err = miniJvm.Start()
	if nil != err {
		t.Fatal(err)
	}

}
