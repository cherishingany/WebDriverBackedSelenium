package control

import (
	"awesomeProject2/service"
	"awesomeProject2/until"
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func Assets(){

	driver := service.RemoteDriver()

	defer driver.Close()

	Login(driver)

	until.Windows(driver)


	elem, _ := driver.FindElement(selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button")
	//
	elem.Click()

	time.Sleep(3 * time.Second)


	elem, _ = driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[2]/form/div/div/div[2]/div/div[2]/div[1]/div/div[2]/div/div[2]/div/div[2]/div/div/div/div/span/span/span/span\n")

	elem.Click()
	time.Sleep(3 * time.Second)
	//elem, err = driver.FindElement(selenium.ByXPATH, "\t/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[31]/div[2]/div/div[2]/div[2]/div/form/div/div/div/div/div/div/table/tbody/tr[2]/td[1]/a\n")
	//elem.Click()

	//_, _ = driver.FindElement(selenium.ByCSSSelector, "div[class=\"ant-table-tbody\"]")

	// 获取WebElements下的内容
	WebElements, _ := driver.FindElements(selenium.ByCSSSelector, "#root > div > section > section > main > div.antd-pro-pages-app-render-index-app > div > div > div > div:nth-child(31) > div.ant-modal-wrap > div > div.ant-modal-content.react-draggable > div.ant-modal-body > div > form > div > div > div > div > div > div > table > tbody > tr:nth-child(n)")

	//	js2 := "var q=document.getElementsByClassName('ant-dropdown-menu ant-dropdown-menu-light ant-dropdown-menu-root ant-dropdown-menu-vertical').li[0].click()"
	//driver.ExecuteScriptRaw(js2,nil)

	for _, values := range WebElements {
		s, _ := values.Text()
		fmt.Println(s)
	}

	time.Sleep(3 * time.Second)
	//	int size = driver.findElement(By.cssSelector(“#formTable > tbody”)).findElements(By.tagName(“tr”)).size();
	//  driver.FindElements(selenium.ByXPATH,"/html/body/div[1]/div/section/section/main/div[2]/div/div/div/div[31]/div[2]/div/div[2]/div[2]/div/form/div/div/div/div/div/div/table/tbody" )


}
