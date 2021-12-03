package control

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func UserLogin(driver selenium.WebDriver) {

	//打开登录页
	err := driver.Get("http://stage.raiborn:8005/")
	if err != nil {
		fmt.Println("you connect error:  ", err.Error())
		return
	}

}
