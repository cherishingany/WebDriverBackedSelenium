package control

import (
	"demo/until"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"time"
)

const (
	seleniumPath = `C:\Users\25834\AppData\Local\Google\Chrome\Application\chromedriver.exe`
	port         = 9515
)

func Selenium1() {

	//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
	var opts []selenium.ServiceOption
	//opts := []selenium.ServiceOption{
	//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
	//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
	//}

	until.SetDebug(true)
	service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if nil != err {
		fmt.Println("start a chromedriver service falid", err.Error())
		return
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
	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("connect to the webDriver faild", err.Error())
		return
	}
	defer driver.Close()
	//打开一个网页
	err = driver.Get("http://stage.raiborn:8005/")
	if err != nil {
		fmt.Println("you connect error:  ", err.Error())
		return
	}
	//

	elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"loginId\"]")

	_ = elem.SendKeys("zwq")

	elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"password\"]")

	_ = elem.SendKeys("1")

	elem, err = driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/div/div[2]/div/div[2]/div[2]/button")

	if nil != err {
		fmt.Printf("search button not found, err: %s", err)
	}
	elem.Click()

	time.Sleep(3 * time.Second)

	//err = driver.Get("http://stage.raiborn:8005/admin/person")
	//if err != nil {
	//	fmt.Println("you connect error:  ", err.Error())
	//	return
	//}
	//
	//_ = driver.AcceptAlert()
	//
	//time.Sleep(3 * time.Second)
	//
	//elems, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[1]")
	//
	//elems.Click()
	//
	//until.DebugLog("hahaha")
	//
	////return
	//time.Sleep(5 * time.Second)
	//
	//elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"person_tabgroup_21_tab_2_sectionrow_138_sectioncol_149_textbox_1712\"]")
	//
	//_ = elem.SendKeys("211")
	//
	//elem, err = driver.FindElement(selenium.ByID, "person_tabgroup_21_tab_2_sectionrow_138_sectioncol_149_textbox_2015")
	//
	//_ = elem.SendKeys("211")
	//
	//time.Sleep(5 * time.Second)
	//
	//elems, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[3]")
	//
	//elems.Click()
	//
	//time.Sleep(5 * time.Second)

	//弹框事件

	err = driver.Get("http://stage.raiborn:8005/assets/asset")
	if err != nil {
		fmt.Println("you connect error:  ", err.Error())
		return
	}

	_ = driver.AcceptAlert()

	time.Sleep(3 * time.Second)
	//

	elems, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button")
	//
	elems.Click()

	time.Sleep(3 * time.Second)

	//分类
	until.ClickEvent(driver, selenium.ByXPATH, "//*[@id=\"rc-tabs-0-panel-canvas1606645094414_tabgroup_62_tab_2\"]/div[1]/div/div[1]/div/div[5]/div/div[2]/div/div/div[1]/div/div/div/span/span/span/span/span")

	elems, err = driver.FindElement(selenium.ByXPATH, "/html/body/div[7]/div/div/ul/li[1]/span")

	//
	elems.Click()

	time.Sleep(3 * time.Second)
}
