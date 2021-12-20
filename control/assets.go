package control

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
	"webdriverbackedselenium/menus"
	"webdriverbackedselenium/util"
)

func Assets(driver selenium.WebDriver) {
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

	WebElement, _ := driver.FindElements(selenium.ByCSSSelector, "div.ant-form-item-control-input-content >span >input")

	for index, values := range WebElement {

		if index < 2 {
			values.SendKeys(fmt.Sprintf("selenium - %d", time.Now().UnixNano()))
		}
	}

	elem, _ = driver.FindElement(selenium.ByCSSSelector, "#rc-tabs-0-panel-canvas1606645094414_tabgroup_62_tab_1 > div > div.ant-table-wrapper > div > div > div > div.ant-table-container > div > table > tbody > tr:nth-child(1) > td:nth-child(1) > span")
	//
	elem.Click()

	time.Sleep(3 * time.Second)

	elem, _ = driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[2]/form/div/div/div[2]/div/div[2]/div[1]/div/div[2]/div/div[2]/div/div[2]/div/div/div/div/span/span/span/span\n")

	elem.Click()
	time.Sleep(3 * time.Second)

	s2 := "#root > div > section > section > main > div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(31) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-body > div > form > div > div > div > div > div > div > table > tbody > tr > td > a"

	elems, _ := driver.FindElements(selenium.ByCSSSelector, s2)

	util.RandomValues(elems)

	//分类
	util.ClickEvent(driver, selenium.ByXPATH, "//*[@id=\"rc-tabs-0-panel-canvas1606645094414_tabgroup_62_tab_2\"]/div[1]/div/div[1]/div/div[5]/div/div[2]/div/div/div[1]/div/div/div/span/span/span/span/span")
	util.ClickEvent(driver, selenium.ByCSSSelector, "li.ant-dropdown-menu-item")
	util.TimeSleep()

	WebElement, _ = driver.FindElements(selenium.ByCSSSelector, " span.ant-tree-node-content-wrapper.ant-tree-node-content-wrapper-close > span.ant-tree-title")

	for index, _ := range WebElement {
		if index == 2 {
			WebElement[2].Click()
			util.TimeSleep()
		}

	}

	//util.RandomValues(WebElement)
	//变更状态 - BUG
	//s1 := "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]"
	//
	//elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	//elem.Click()
	//
	//util.TimeSleep()
	//
	//s1 = "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]\n"
	//
	//elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	//elem.Click()
	//
	//util.TimeSleep()
	//util.TimeSleep()
	////第一层div   "ant-select-show-search"
	//s2 = "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[13]/div[2]/div/div[2]/div[2]/form/div[1]/div/div[2]/div/div/div"
	//elem, _ = driver.FindElement(selenium.ByXPATH, s2)
	//elem.Click()
	//
	//s3 := "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[13]/div[2]/div/div[2]/div[2]/form/div[1]/div/div[2]/div/div/div/span[1]/span"
	//elem, _ = driver.FindElement(selenium.ByXPATH, s3)
	//elem.Click()
	//
	//time.Sleep(time.Second)
	//
	////从下拉框中选择相应的活动事件

	//selenium.WebDriver.Wait()

	//操作菜单按钮
	ws, _ := driver.FindElements(selenium.ByCSSSelector, "div.index_appButtons__1ByKw button")
	for index, values := range ws {
		if index == menus.Save {
			values.Click()
			util.TimeSleep()
		}

	}
	//操作按钮 span.ant-input-group-addon > span
	//input 输入框 div.ant-form-item-control-input-content > span > input
	//三角选择框 anticon anticon-unordered-list ant-dropdown-trigger

}
