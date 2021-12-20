package control

import (
	"github.com/tebeka/selenium"
	"log"
	"webdriverbackedselenium/menus"
	"webdriverbackedselenium/model"
	"webdriverbackedselenium/util"
)

func ChooseFile(driver selenium.WebDriver) {

	defer func() {
		driver.Close()
		//wg.Done()
	}()
	Login(driver)
	util.Windows(driver)
	util.TimeSleep()

	//进入增加资产
	elem, _ := driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button")
	//
	elem.Click()

	util.TimeSleep()

	WebElement, _ := driver.FindElements(selenium.ByCSSSelector, "div.index_appContent__1KRQE > form > div > div > div.ant-tabs-nav > div.ant-tabs-nav-wrap > div>div.ant-tabs-tab >div")
	util.TimeSleep()

	asset := model.Asset{}
	for i, v := range WebElement {

		switch i {
		case 1:
			//v.Click()
			asset.Info(driver)
		case 2:
			v.Click()
		case 3:
			v.Click()
		case 4:
			v.Click()
		case 5:
			v.Click()
		case 6:
			v.Click()
		default:
			log.Println("lens error -----------------------------")
		}

	}

	//操作菜单按钮
	ws, _ := driver.FindElements(selenium.ByCSSSelector, "div.index_appButtons__1ByKw button")
	for index, values := range ws {
		if index == menus.Save {
			values.Click()
			util.TimeSleep()
		}

	}

	util.TimeSleep()
	//变更资产状态
	asset.ChooseStatus(driver)

}
