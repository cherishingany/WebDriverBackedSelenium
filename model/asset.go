package model

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
	"webdriverbackedselenium/util"
)

type Asset struct {
}

func (*Asset) Info(driver selenium.WebDriver) {
	WebElement, _ := driver.FindElements(selenium.ByCSSSelector, "div.ant-form-item-control-input-content >span >input")

	for index, values := range WebElement {

		if index < 2 {
			values.SendKeys(fmt.Sprintf("selenium - %d", time.Now().UnixNano()))
		}
	}

	elem, _ := driver.FindElement(selenium.ByCSSSelector, "#rc-tabs-0-panel-canvas1606645094414_tabgroup_62_tab_1 > div > div.ant-table-wrapper > div > div > div > div.ant-table-container > div > table > tbody > tr:nth-child(1) > td:nth-child(1) > span")
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
}

func (*Asset) ChooseStatus(driver selenium.WebDriver) {

	//变更状态 - BUG
	//defer func() {
	//	driver.Close()
	//	//wg.Done()
	//}()
	//
	//Login(driver)
	//util.Windows(driver)
	//util.TimeSleep()

	s := "div > div.ant-table-container > div > table > tbody > tr:nth-child(10) > td:nth-child(1) > span >a"

	elem, _ := driver.FindElement(selenium.ByCSSSelector, s)
	elem.Click()
	util.TimeSleep()

	//变更状态 - BUG
	s1 := "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]"

	elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	elem.Click()

	util.TimeSleep()

	s1 = "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]\n"

	elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	elem.Click()

	util.TimeSleep()
	//第一层div   "ant-select-show-search"
	s1 = "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[13]/div[2]/div/div[2]/div[2]/form/div[1]/div/div[2]/div/div/div"
	elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	elem.Click()
	util.TimeSleep()
	s3 := "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[13]/div[2]/div/div[2]/div[2]/form/div[1]/div/div[2]/div/div/div/span[1]/span"
	elem, _ = driver.FindElement(selenium.ByXPATH, s3)
	elem.Click()

	s4 := "div.rc-virtual-list-holder-inner > div >div"
	elems, _ := driver.FindElements(selenium.ByCSSSelector, s4)
	for i, v := range elems {
		if i == 6 {
			v.Click()
		}
	}
	util.TimeSleep()
	//到未就绪状态
	s5 := "div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(13) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-footer > button:nth-child(1)"

	elem, _ = driver.FindElement(selenium.ByCSSSelector, s5)

	elem.Click()
	util.TimeSleep()

	//操作按钮 span.ant-input-group-addon > span
	//input 输入框 div.ant-form-item-control-input-content > span > input
	//三角选择框 anticon anticon-unordered-list ant-dropdown-trigger
}
