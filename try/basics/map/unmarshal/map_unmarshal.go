package unmarshal

import (
	"encoding/json"
	"log"
	"time"
)

type people struct {
	Name     string     `json:"name"`
	City     string     `json:"city"`
	Birthday *time.Time `json:"birthday"`
}

// Unmarshal Unmarshal
func Unmarshal() (err error) {
	data := map[string]string{
		"name":     "123",
		"city":     "234",
		"birthday": "",
	}
	if data["birthday"] == "" {
		delete(data, "birthday")
	}
	itemJOSN, err := json.Marshal(data)
	if err != nil {
		log.Printf("Unmarshal data:%v,err:%v", data, err)
		return
	}
	p1 := &people{}
	err = json.Unmarshal(itemJOSN, p1)
	if err != nil {
		log.Printf("Unmarshal data:%v,err:%v", data, err)
		return
	}
	return
}
