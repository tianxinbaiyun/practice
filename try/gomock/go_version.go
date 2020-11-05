package main

import (
	"github.com/tianxinbaiyun/practice/try/gomock/spider"
)

// GetGoVersion GetGoVersion
func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
