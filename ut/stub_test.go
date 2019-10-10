package main

import (
	"testing"
)

func TestHandle(t *testing.T) {
	res := Handle(1, 3, 0)
	t.Log(res, 1, 3)
	res = Handle(2, 5, 1)
	t.Log(res, 2, 5)
	res = Handle(5, 1, -2)
	t.Log(res, 5, 1)
	
}

// Stub 操作，用桩函数替换掉真是的GetType函数
func StubFunc(a int) string {
	// 不管a的值是多少，我们都输出"add"
	return "add" 
}
func TestHandleStub(t *testing.T) {
	CallFunc = StubFunc       
	res := Handle(1, 3, 0)
	t.Log(res, 1, 3)
	res = Handle(2, 5, 1)
	t.Log(res, 2, 5)
	res = Handle(5, 1, -2)
	t.Log(res, 5, 1)
	
}
