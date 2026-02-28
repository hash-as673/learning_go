package main;

import "fmt"

func validateConfig(envName string, portNo int) (message string,isValid bool) {
	switch (envName) {
	case "prod":
		if portNo == 443 || portNo == 80{
			message = "Valid Prod Config"
			isValid = true
			} else {
				message = "Invalid Port for Prod"
				isValid = false
			}
		case "dev":
			if portNo > 1024 && portNo < 65535{
				message = "Valid Dev Config"
			isValid = true
			
			} else {
				message = "Invalid Port for Dev"
				isValid = false
				
			}
		default:
			message = "Invalid Config"
			isValid = false
	}
	return
}

func main(){
	 fmt.Println("Case 1: Valid Prod:")
	 envName := "prod"
	 portNo := 443
	 fmt.Println(validateConfig(envName,portNo))
	 fmt.Println("Case 2: Invalid Prod Port No:")
	 portNo = 440
	 fmt.Println(validateConfig(envName,portNo))
	 fmt.Println("Case 3: Valid Dev:")
	 envName = "dev"
	 portNo = 1025
	 fmt.Println(validateConfig(envName,portNo))
	
}