package ticker

import (
	"log"
	"time"
)

func regisiter() {

	// get self ip
	tk := time.NewTicker(13 * time.Second)
	for {
		log.Println("111")
		<-tk.C
	}
}
