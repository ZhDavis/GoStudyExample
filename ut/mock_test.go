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

// Mock 的操作方式, mock一个struct实现interface里面的函数
type MockService struct {
}

func (s *MockService)GetType(flag int) string {
	return "add"
}

func TestHandleMock(t *testing.T) {
        client := &MockService{}
	res := Handle(1, 3, 0, client)
	t.Log(res, 1, 3)
	res = Handle(2, 5, 1, client)
	t.Log(res, 2, 5)
	res = Handle(5, 1, -2, client)
	t.Log(res, 5, 1)
}
