package main

type Client interface {
    GetType(flag int) string
}

type Service struct {
}

func (c *Service) GetType(flag int) string {
    if flag > 0 {
        return "add"    
    } 
    if flag < 0 {
        return "sub"
    }
    return "zero"
}

func Handle(a,b, flag int, client Client) int {
    switch client.GetType(flag) {
        case "add":
            return a + b
        case "sub":
            return a - b
        default:
           return 0
    }
    return 0
}
