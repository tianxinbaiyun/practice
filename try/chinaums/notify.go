package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// ChinaumsNotyfyHeader ChinaumsNotyfyHeader
type ChinaumsNotyfyHeader struct {
	Version     string `json:"version"`
	Transtype   string `json:"transtype"`
	Employno    string `json:"employno"`
	Termid      string `json:"termid"`
	RequestTime string `json:"request_time"`
}

// ChinaumsNotyfyBody ChinaumsNotyfyBody
type ChinaumsNotyfyBody struct {
	Orderno   string `json:"orderno"`
	Cod       string `json:"cod"`
	Payway    string `json:"payway"`
	Banktrace string `json:"banktrace"`
	Postrace  string `json:"postrace"`
	Tracetime string `json:"tracetime"`
	Cardid    string `json:"cardid"`
	Signflag  string `json:"signflag"`
	Signer    string `json:"signer"`
	QueryID   string `json:"queryID"`
	Memo      string `json:"memo"`
}

// ChinaumsNotyfyRsp ChinaumsNotyfyRsp
type ChinaumsNotyfyRsp struct {
	ChinaumsNotyfyHeader ChinaumsNotyfyHeader `json:"header"`
	ChinaumsNotyfyBody   ChinaumsNotyfyBody   `json:"body"`
}

func main() {
	str := `context={"header":{"version":"1.0","transtype":"P033","employno":"01","termid":"12345678","request_time":"20120913140101"},"body":{"orderno":"041233939394","cod":"100.88","payway":"02","banktrace":"123456789012","cardid":"390284793029340295"}}&mac=53B983CE7535E80604F7A7A33C03F099`
	notify(str)
}
func notify(str string) (err error) {
	m, _ := url.ParseQuery(str)
	for i, i2 := range m {
		fmt.Println(i, ":", i2)
	}
	fmt.Println(m["context"][0])
	fmt.Println(m["mac"][0])
	signKey := "1111111111111111111111111111111111111111111111111111111111111111"
	preSignStr := fmt.Sprintf("%s%s", m["context"][0], signKey)
	fmt.Println(preSignStr)
	sign := SignMd5(preSignStr)
	fmt.Println(sign)
	sign = strings.ToUpper(sign)
	if sign == m["mac"][0] {
		fmt.Println(111)
	}
	orderRsp := &ChinaumsNotyfyRsp{}
	err = json.Unmarshal([]byte(m["context"][0]), orderRsp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(orderRsp)

	return
}
