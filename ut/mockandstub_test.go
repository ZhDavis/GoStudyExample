package main

import "testing"

func TestHandle(t *testing.T) {
        client := &Service{}
	res := Handle(1, 3, 0, client)
	t.Log(res, 1, 3)
	res = Handle(2, 5, 1, client)
	t.Log(res, 2, 5)
	res = Handle(5, 1, -2, client)
	t.Log(res, 5, 1)
}

// Mock和stub结合的操作方式, mock一个struct实现interface里面的函数
// 这个函数采用回调函数的方式来处理
type MockService struct {
	StubFunc func(flag int) string
        StubFuncUsed bool
}

func (s *MockService)GetType(flag int) string {
        s.StubFuncUsed = true
	return s.StubFunc(flag)
}

func TestHandleMock(t *testing.T) {
        client := &MockService{}
        client.StubFunc = func(flag int) string{
		return "add"
        }
	res := Handle(1, 3, 0, client)
	t.Log(res, 1, 3)
	res = Handle(2, 5, 1, client)
	t.Log(res, 2, 5)
        client.StubFunc = func(flag int) string{
		return "zero"
        }
	res = Handle(5, 1, -2, client)
	t.Log(res, 5, 1)
}
