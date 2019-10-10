package main

import (
	"testing"
)

// 1. 定义测试校验的数据结构
type TestTable struct {
	inArg  int // 测试入参
	result int // 测试结果 
}

func TestAbs(t *testing.T) {
	// 2. 构造的测试结果
	tables := []TestTable {
		{0, 0},
		{5, 5},
		{-5, 5},
	}
	// 3. 函数测试校验
	for _,test := range tables {
		res := Abs(test.inArg)
		t.Log(res, test.result)
	}
}
