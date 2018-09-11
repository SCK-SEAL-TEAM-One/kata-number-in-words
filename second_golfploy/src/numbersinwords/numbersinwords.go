package numbersinwords

import "strconv"

func ConvertNumberToThaiCharactor(inputNumber float64) string {
	number := int(inputNumber)
	stringNumber := strconv.Itoa(number)
	unit := len(stringNumber)
	var numberThai string
	var unitThai string
	for _, numberCharactor := range stringNumber {
		if string(numberCharactor) == "1" {
			numberThai = "หนึ่ง"
			if unit == 6 {
				unitThai = "แสนบาทถ้วน"
			}
		}
	}
	return numberThai + unitThai
}
