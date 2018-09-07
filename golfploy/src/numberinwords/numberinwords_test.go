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
