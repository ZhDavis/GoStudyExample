package  main

// 计算输入参数的绝对值
func Abs(a int) (result uint32) {
	if a == 0 {
		return
        }

	if a > 0 {
		result = uint32(a)
		return 
	}
	
	if a < 0 {
		result = uint32(-a)
		return
	}
	return
}
