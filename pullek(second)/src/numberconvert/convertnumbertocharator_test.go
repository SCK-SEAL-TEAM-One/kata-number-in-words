package numberconvert_test

import (
	"numberconvert"
	"testing"
)

func Test_NumberConvert_Input_100000_Should_Be_One_Hundred_Thousand_Baht(t *testing.T) {
	expectedNumber := "หนึ่งแสนบาทถ้วน"
	number := 100000

	actual := numberconvert.NumberConvert(number)

	if expectedNumber != actual {
		t.Errorf("expect %s but it got %s", expectedNumber, actual)
	}
}
func Test_NumberConvert_Input_500000_Should_Be_Five_Hundred_Thousand_Baht(t *testing.T) {
	expectedNumber := "ห้าแสนบาทถ้วน"
	number := 500000

	actual := numberconvert.NumberConvert(number)

	if expectedNumber != actual {
		t.Errorf("expect %s but it got %s", expectedNumber, actual)
	}
}
func Test_NumberConvert_Input_50000_Should_Be_Fifty_Thousand_Baht(t *testing.T) {
	expectedNumber := "ห้าหมื่นบาทถ้วน"
	number := 50000

	actual := numberconvert.NumberConvert(number)

	if expectedNumber != actual {
		t.Errorf("expect %s but it got %s", expectedNumber, actual)
	}
}
