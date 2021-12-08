package until

import (
	"fmt"
	"github.com/tebeka/selenium"
	"time"
)

func TimeSleep() {
	time.Sleep(3 * time.Second)
}

func Windows(driver selenium.WebDriver) {

	e := driver.Get("http://stage.raiborn:8005/assets/asset")
	if e != nil {
		fmt.Println("you connect error:  ", e.Error())
		return
	}

	//弹框事件
	_ = driver.AcceptAlert()

	time.Sleep(3 * time.Second)

}

func ClickEvent(driver selenium.WebDriver, by, value string) {
	elem, err := driver.FindElement(by, value)
	if err != nil {
		panic(err)

	}
	err = elem.Click()
	if err != nil {
		panic(err)
	}
	time.Sleep(3 * time.Second)

}