package main

import (
	"fmt"
)

type testiface interface {
	SayHello()
	Say(s string)
	Increment()
	GetInternalValue() int
}

type testConcreteImpl struct {
	i int
}

func NewTestConcreteImpl() testiface { // this is our constructor...
	return new(testConcreteImpl) //&testConcreteImpl{}
}

func NewTestConcreteImplWithV(v int) testiface {
	return &testConcreteImpl{i: v}
}

func (tst *testConcreteImpl) SayHello() {
	fmt.Println("Hello")
}

func (tst *testConcreteImpl) Say(s string) {
	fmt.Println(s)
}

func (tst *testConcreteImpl) Increment() {
	tst.i++
}

func (tst *testConcreteImpl) GetInternalValue() int {
	return tst.i
}

type testEmbedding struct { //we want this struct to have all the features of *testConcreteImpl, this is called the outer type
	*testConcreteImpl //embedding, this is called the inner type
}

func testEIface(v interface{}) {
	/*
		if i, ok := v.(int); ok { //type assertion
			fmt.Println("I am an integer my value is: ", i)
		} else {
			fmt.Println("I am not an integer type!!")
		}
	*/

	switch val := v.(type) { //type switch
	case int:
		fmt.Println("I am an int ", val)
	case string:
		fmt.Println("I am a string ", val)
	default:
		fmt.Println("I am neither an int nor a string ", val)
	}
}

func main() {
	var tiface testiface
	tiface = NewTestConcreteImplWithV(5) //&testConcreteImpl{} //new(testConcreteImpl)
	tiface.SayHello()
	tiface.Say("Hello again!!")
	tiface.Increment()
	tiface.Increment()
	tiface.Increment()
	fmt.Println(tiface.GetInternalValue())
	te := testEmbedding{testConcreteImpl: &testConcreteImpl{i: 50}}
	te.SayHello() //te.testConcreteImpl.SayHello()
	te.Increment()
	fmt.Println(te.GetInternalValue())
	testEIface(3)
	testEIface("string to empty interface")
	testEIface(tiface)
}
