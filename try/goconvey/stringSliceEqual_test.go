package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//func TestStringSliceEqual(t *testing.T) {
//	Convey("TestStringSliceEqual should return true when a != nil  && b != nil", t, func() {
//		a := []string{"hello", "goconvey"}
//		b := []string{"hello", "goconvey"}
//		So(StringSliceEqual(a, b), ShouldBeTrue)
//	})
//}

func TestStringSliceEqual(t *testing.T) {
	Convey("TestStringSliceEqual should return true when a!=nil && b!=nil", t, func() {
		a := []string{"hello", "goconvey"}
		b := []string{"hello", "goconvey", "123"}
		So(StringSliceEqual(a, b), ShouldBeFalse)
	})
}
