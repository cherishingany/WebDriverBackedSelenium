package control

import (
	"github.com/tebeka/selenium"
	"time"
	"webdriverbackedselenium/until"
)

func Assets(driver selenium.WebDriver) {
	//wg *sync.WaitGroup

	defer func() {
		driver.Close()
		//wg.Done()
	}()

	Login(driver)

	until.Windows(driver)

	//增加资产
	//elem, _ := driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button")
	////
	//elem.Click()

	elem, _ := driver.FindElement(selenium.ByCSSSelector, "#rc-tabs-0-panel-canvas1606645094414_tabgroup_62_tab_1 > div > div.ant-table-wrapper > div > div > div > div.ant-table-container > div > table > tbody > tr:nth-child(1) > td:nth-child(1) > span")
	//
	elem.Click()

	time.Sleep(3 * time.Second)

	elem, _ = driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[2]/form/div/div/div[2]/div/div[2]/div[1]/div/div[2]/div/div[2]/div/div[2]/div/div/div/div/span/span/span/span\n")

	elem.Click()
	time.Sleep(3 * time.Second)

	s2 := "#root > div > section > section > main > div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(31) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-body > div > form > div > div > div > div > div > div > table > tbody > tr > td > a"

	elems, _ := driver.FindElements(selenium.ByCSSSelector, s2)

	until.RandomValues(elems)

	//for index, values := range elems {
	//	s, _ := values.Text()
	//	rand.Seed(time.Now().UnixNano())
	//	x := rand.Intn(len(elems))
	//	fmt.Println("此次生成的随机数为: ", x)
	//	if x == index {
	//		values.Click()
	//	}
	//	fmt.Println(index, s)
	//}

	//elem, err = driver.FindElement(selenium.ByXPATH, "\t/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[31]/div[2]/div/div[2]/div[2]/div/form/div/div/div/div/div/div/table/tbody/tr[2]/td[1]/a\n")
	//elem.Click()

	//_, _ = driver.FindElement(selenium.ByCSSSelector, "div[class=\"ant-table-tbody\"]")

	// 获取WebElements下的内容
	//WebElements, _ := driver.FindElements(selenium.ByCSSSelector, "#root > div > section > section > main > div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(31) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-body > div > form > div > div > div > div > div > div > table > tbody > tr:nth-child(n)")

	//	js2 := "var q=document.getElementsByClassName('ant-dropdown-menu ant-dropdown-menu-light ant-dropdown-menu-root ant-dropdown-menu-vertical').li[0].click()"
	//driver.ExecuteScriptRaw(js2,nil)

	//for _, values := range WebElements {
	//	s, _ := values.Text()
	//	fmt.Println(s)
	//}

	time.Sleep(3 * time.Second)
	//	int size = driver.findElement(By.cssSelector(“#formTable > tbody”)).findElements(By.tagName(“tr”)).size();
	//  driver.FindElements(selenium.ByXPATH,"/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[31]/div[2]/div/div[2]/div[2]/div/form/div/div/div/div/div/div/table/tbody" )

	//变更状态
	//s1 := "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]"
	//
	//elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	//elem.Click()
	//
	//until.TimeSleep()
	//
	//s1 = "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[6]\n"
	//
	//elem, _ = driver.FindElement(selenium.ByXPATH, s1)
	//elem.Click()
	//
	//until.TimeSleep()
	//until.TimeSleep()
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
	//s1 = "div.rc-virtual-list"
	//elem, _ = driver.FindElement(selenium.ByCSSSelector, s1)
	//elem.Click()
	//
	//time.Sleep(time.Second)
	//
	//elem, _ = driver.FindElement(selenium.ByCSSSelector, "div.rc-virtual-list > div.rc-virtual-list-holder > div > div > div:nth-child(1) > span")
	//elem.Click()
	//
	//time.Sleep(time.Second)

	//selenium.WebDriver.Wait()

	//操作菜单按钮
	ws, _ := driver.FindElements(selenium.ByCSSSelector, "div.index_appButtons__1ByKw button")
	for index, values := range ws {
		if index == 6 {
			values.Click()
		}

		until.TimeSleep()
	}

}
