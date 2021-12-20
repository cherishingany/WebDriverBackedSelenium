package util

import (
	"fmt"
	"github.com/tebeka/selenium"
	"math/rand"
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
	driver.AcceptAlert()

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
	time.Sleep(2 * time.Second)

}

func Click(driver selenium.WebDriver) {

	//鼠标悬停
	//above, _ := driver.FindElement(selenium.ByCSSSelector, "tj_settingicon") //通过name找到设置按钮
	//above.MoveTo()
	//ActionChains(driver).move_to_element(above).perform() #move_to_element移到设置的元素, avove上面定位到的设置.然后执行操作
	//sleep(4)

}

func RandomValues(WebElements []selenium.WebElement) {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(len(WebElements))
	fmt.Println("此次生成的随机数为: ", x)

	WebElements[x].Click()
	TimeSleep()
}
