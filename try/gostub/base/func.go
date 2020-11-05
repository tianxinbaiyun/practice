package base

import (
	"time"

	"fmt"

	"os"

	"github.com/prashantv/gostub"
)

var timeNow = time.Now
var osHostname = os.Hostname

func getDate() int {
	return timeNow().Day()
}
func getHostName() (string, error) {
	return osHostname()
}

func StubTimeNowFunction() {
	stubs := gostub.Stub(&timeNow, func() time.Time {
		return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
	})
	fmt.Println(getDate())
	defer stubs.Reset()
}

func StubHostNameFunction() {
	stubs := gostub.StubFunc(&osHostname, "LocalHost", nil)
	defer stubs.Reset()
	fmt.Println(getHostName())
}

func StubTimeNowFunction2() int {
	stubs := gostub.Stub(&timeNow, func() time.Time {
		return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
	})
	fmt.Println(getDate())
	defer stubs.Reset()
	return getDate()
}
