package main

import (
	"testing"  // 1.必须要包含UT的单元测试包testing
)
// 2.测试函数命名必须是Test开头，一般需要Test+要测试的函数名
// 3.参数一定要包含*testing.T 这一个入参
func TestAdd(t *testing.T) { 
	t.Log(Add(1), 2)
}

