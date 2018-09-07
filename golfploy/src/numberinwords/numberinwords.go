package numberinwords

import (
	"fmt"
)

func ConvertDecimalToThaiCharactor(amountMoney float64) string {
	stringAmountMoney := fmt.Sprintf("%d", int(amountMoney))
	unit := len(stringAmountMoney)
	for _, runeValue := range stringAmountMoney {
		if unit == 6 && string(runeValue) == "1" {
			thaiNumber := "หนึ่งแสน"
			return thaiNumber + "บาทถ้วน"
		}
	}

	return ""
}
