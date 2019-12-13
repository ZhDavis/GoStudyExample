package main


func HandleType(flag int) string {
	switch flag {
		case 0:
			return "Add"
                case 1:
                        return "Sub"
                case 2: 
                        return "Max"
        }
        return ""
}
