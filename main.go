package main

import (
	"awesomeProject2/control"
	"awesomeProject2/service"
	"fmt"

)

func main() {
	fmt.Println("hello go ")
	startWebDriver := service.StartWebDriver()
	defer startWebDriver.Stop()

	control.Assets()

	//	group := sync.WaitGroup{}
	//	group.Add(1)
	//
	//	func() {
	//		go control.Assets()
	//		group.Wait()
	//	}()
	//
	//	group.Done()
	//
}