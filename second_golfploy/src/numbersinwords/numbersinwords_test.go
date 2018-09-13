package numbersinwords_test

import (
	. "numbersinwords"
	"testing"
)

func Test_ConvertNumberToThaiCharactor_Input_100000_Should_Be_OneHundredThousand(t *testing.T) {
	expectedResult := "หนึ่งแสนบาทถ้วน"
	inputNumber := 100000.00

	actualResult := ConvertNumberToThaiCharactor(inputNumber)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}

func Test_ConvertNumberToThaiCharactor_Input_400000_Should_Be_Four_Hundred_Thousand_Baht(t *testing.T) {
	expectedResult := "สี่แสนบาทถ้วน"
	inputNumber := 400000.00

	actualResult := ConvertNumberToThaiCharactor(inputNumber)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}

func Test_ConvertNumberToThaiCharactor_Input_50000_Should_Be_Fifty_Thousand_Baht(t *testing.T) {
	expectedResult := "ห้าหมื่นบาทถ้วน"
	inputNumber := 50000.00

	actualResult := ConvertNumberToThaiCharactor(inputNumber)

	if expectedResult != actualResult {
		t.Errorf("expect %s but it got %s", expectedResult, actualResult)
	}
}
