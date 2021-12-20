package service

import (
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	SeleniumPath = `C:\Users\25834\AppData\Local\Google\Chrome\Application\chromedriver.exe`
	PORT         = 9515
)

func StartWebDriver() *selenium.Service {

	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	var opts []selenium.ServiceOption
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}
	//selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(SeleniumPath, PORT, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		panic(err)
	}
	//注意这里，server关闭之后，chrome窗口也会关闭
	//关闭一个webDriver会对应关闭一个chrome窗口
	//但是不会导致seleniumServer关闭
	return service
}
