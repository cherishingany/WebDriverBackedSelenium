package control

import (
	"demo/until"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	SeleniumPath = `C:\Users\25834\AppData\Local\Google\Chrome\Application\chromedriver.exe`
	PORT         = 9515
)

func DriverService() selenium.WebDriver {

	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	var opts []selenium.ServiceOption
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}

	until.SetDebug(true)
	service, err := selenium.NewChromeDriverService(SeleniumPath, PORT, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		panic(err)
	}
	//注意这里，server关闭之后，chrome窗口也会关闭
	defer service.Stop()

	//链接本地的浏览器 chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	//禁止图片加载，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
			//"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)

	// 重新调起chrome浏览器
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", PORT))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		panic(err)
	}
	defer driver.Close()

	return driver

}
