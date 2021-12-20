package main

import (
	"fmt"
	"webdriverbackedselenium/control"
	"webdriverbackedselenium/service"
	"webdriverbackedselenium/util"
)

func main() {
	fmt.Println("hello go ")
	startWebDriver := service.StartWebDriver()

	defer startWebDriver.Stop()
	driver := service.RemoteDriver()

	util.TimeSleep()

	//control.Login(driver)
	//control.UserRegister(driver)
	//control.Assets(driver)

	control.ChooseFile(driver)

}
