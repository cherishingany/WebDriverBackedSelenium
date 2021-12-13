package main

import (
	"awesomeProject2/control"
	"awesomeProject2/service"
	"awesomeProject2/until"
	"fmt"
)

func main() {
	fmt.Println("hello go ")
	startWebDriver := service.StartWebDriver()

	defer startWebDriver.Stop()
	driver := service.RemoteDriver()

	until.TimeSleep()
	//control.Assets(driver)

	control.Login(driver)
}
