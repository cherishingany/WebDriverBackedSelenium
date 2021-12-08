package service

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func RemoteDriver() selenium.WebDriver {

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

	//调起chrome浏览器
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", PORT))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		panic(err)
	}

	return driver
}
