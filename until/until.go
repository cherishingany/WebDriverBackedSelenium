package until

import "log"

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
