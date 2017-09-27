package main

import (
	"PropertiesScanner/utils"
		"PropertiesScanner/constants"
		"fmt"
)


func main() {
	
	//Reading values from properties file
	
	/*Here I have created ResourceManager struct like we are using
	in java to read the properties file. So I've mapped the constants 
	with the messages from properties file. It is a standard approach to
	read the values from properties file */
	
	resourceManager:=utils.ResourceManager{};
	successMsg:=resourceManager.GetProperty(constants.LOGIN_SUCCESS)
	notFoundMsg:=resourceManager.GetProperty(constants.NOT_FOUND)
	fmt.Println(successMsg)
	fmt.Println(notFoundMsg)
}
