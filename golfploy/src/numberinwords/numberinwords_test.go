package numberinwords_test

import (
	. "numberinwords"
	"testing"
)

func Test_ConvertDecimalToThaiCharactor_Input_100000_Should_Be_One_Hundred_Thousand_Baht(t *testing.T) {
	expectedResult := "หนึ่งแสนบาทถ้วน"
	amountMoney := 100000.00

	actualResult := ConvertDecimalToThaiCharactor(amountMoney)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}

func Test_ConvertDecimalToThaiCharactor_Input_400000_Should_Be_Four_Hundred_Thousand_Baht(t *testing.T) {
	expectedResult := "สี่แสนบาทถ้วน"
	amountMoney := 400000.00

	actualResult := ConvertDecimalToThaiCharactor(amountMoney)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}

func Test_ConvertDecimalToThaiCharactor_Input_50000_Should_Be_Fifty_Thousand_Baht(t *testing.T) {
	expectedResult := "ห้าหมื่นบาทถ้วน"
	amountMoney := 50000.00

	actualResult := ConvertDecimalToThaiCharactor(amountMoney)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}
