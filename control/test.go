package control

import (
	"github.com/tebeka/selenium"
	"webdriverbackedselenium/util"
)

func Test(driver selenium.WebDriver) {

	defer func() {
		driver.Close()
		//wg.Done()
	}()

	Login(driver)
	util.Windows(driver)
	util.TimeSleep()

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

		if i == 7 {
			v.Click()
		}
	}

	//到未就绪状态
	s5 := "div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(13) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-footer > button:nth-child(1)"

	elem, _ = driver.FindElement(selenium.ByCSSSelector, s5)

	elem.Click()
	util.TimeSleep()
}
