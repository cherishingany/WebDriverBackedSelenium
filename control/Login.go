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
	//WebElements, err := driver.FindElements(selenium.ByCSSSelector, "div.ant-form-item-control-input-content > input")
	//for _, values := range WebElements {
	//	s, _ := values.Text()
	//	fmt.Println(s)
	//	values.SendKeys("11")
	//}

	//

	time.Sleep(1 * time.Second)

	//
	//elem.SendKeys("1")
	//
	//elem, err = driver.FindElement(selenium.ByXPATH, "//*[@id=\"password\"]")
	//
	////element.SendKeys(selenium.ControlKey + "t")
	//_ = elem.SendKeys("zwq")
	//
	driver.ExecuteScriptRaw("document.getElementById(\"loginId\").value = \"zwq\"", nil)
	driver.ExecuteScriptRaw("document.getElementById(\"password\").value = \"1\"", nil)
	time.Sleep(1 * time.Second)
	elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"loginId\"]")
	elem.SendKeys(selenium.EnterKey)

}
