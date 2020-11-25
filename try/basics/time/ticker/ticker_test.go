package ticker

import (
	"github.com/google/martian/log"
	"github.com/xs23933/fetch"
	"testing"
)

func TestRegisiter(t *testing.T) {
	regisiter()
}

func TestAA(t *testing.T) {
	APIs := map[string]interface{}{
		// http://172.17.0.1:40203
		"http://172.17.0.1:40205": []string{
			"/api/mws_finance_msean/*",
		},
	}
	if buf, err := fetch.Payload("http://172.17.0.1:88/sign", APIs); err != nil {
		log.Infof("Apisign failed %s %v", string(buf), err)
	} else {
		log.Infof("Apisign success %s", string(buf))
	}
}
