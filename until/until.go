package until

import (
	"github.com/tebeka/selenium"
	"log"
	"time"
)

var debugFlag = false

func SetDebug(debug bool) {
	debugFlag = debug
}

func DebugLog(format string, args ...interface{}) {
	if !debugFlag {
		return
	}
	log.Printf("   "+"\n"+format+"\n", args...)
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
