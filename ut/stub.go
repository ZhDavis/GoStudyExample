package main

func GetType(flag int)  string {
	t := ""
	if flag<0 {
		t = "sub"
	} else if flag==0 {
		t = "zero"
	} else {
		t = "add"
	}
	return t
}
 
var CallFunc func(int)(string) = GetType

func Handle(a, b, t int) (result int){
	switch CallFunc(t) {
		case "sub":
			result = a -b
		case "add":
			result = a + b
		case "zero":
			result = 0
	}	
	return
}
