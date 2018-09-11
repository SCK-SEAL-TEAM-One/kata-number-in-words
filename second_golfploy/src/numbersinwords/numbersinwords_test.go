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
