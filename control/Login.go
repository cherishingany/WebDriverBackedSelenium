package control

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func Login(driver selenium.WebDriver) {

	err := driver.Get("http://stage.raiborn:8005/")
	if err != nil {
		fmt.Println("you connect error:  ", err.Error())
		return
	}

	elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"loginId\"]")

	_ = elem.SendKeys("zwq")

	elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"password\"]")
	_ = elem.SendKeys("1")

	_ = elem.SendKeys(selenium.EnterKey)

	time.Sleep(3 * time.Second)

}
