package control

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
	"webdriverbackedselenium/service"
	"webdriverbackedselenium/until"
)

func UserRegister() { //wg *sync.WaitGroup

	driver := service.RemoteDriver()

	defer func() {
		driver.Close()
		//wg.Done()
	}()

	Login(driver)

	err := driver.Get("http://stage.raiborn:8005/admin/person")
	if err != nil {
		fmt.Println("you connect error:  ", err.Error())
		return
	}

	_ = driver.AcceptAlert()

	until.ClickEvent(driver, selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[1]")

	elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"person_tabgroup_21_tab_2_sectionrow_138_sectioncol_149_textbox_1712\"]")

	_ = elem.SendKeys("211")

	elem, err = driver.FindElement(selenium.ByID, "person_tabgroup_21_tab_2_sectionrow_138_sectioncol_149_textbox_2015")

	_ = elem.SendKeys("211")

	time.Sleep(5 * time.Second)

	until.ClickEvent(driver, selenium.ByXPATH, "//*[@id=\"root\"]/div/section/section/main/div[2]/div/div/div/div[1]/div[2]/button[3]")
}
