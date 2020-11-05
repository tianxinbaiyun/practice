package base

import (
	"fmt"
	"github.com/prashantv/gostub"
)

// StubGlobalVariable StubGlobalVariable
func StubGlobalVariable(counter int) {
	stubs := gostub.Stub(&counter, 200)
	defer stubs.Reset()
	fmt.Println("Counter:", counter)
}
