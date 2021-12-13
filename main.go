package main

import (
	"fmt"
	"webdriverbackedselenium/control"
	"webdriverbackedselenium/service"
	"webdriverbackedselenium/until"
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
