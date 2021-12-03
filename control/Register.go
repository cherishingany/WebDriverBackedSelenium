package control

import (
	"demo/until"
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func Register() {

	driver := DriverService()
	UserLogin(driver)

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
