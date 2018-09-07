package numberinwords

import (
	"fmt"
)

func ConvertDecimalToThaiCharactor(amountMoney float64) string {
	stringAmountMoney := fmt.Sprintf("%d", int(amountMoney))
	unitMap := map[int]string{
		6: "แสน",
		5: "หมื่น",
		4: "พัน",
		3: "ร้อย",
		2: "สิบ",
		1: "",
	}
	thaiCharactor := map[string]string{
		"1": "หนึ่ง",
		"2": "สอง",
		"3": "สาม",
		"4": "สี่",
		"5": "ห้า",
		"6": "หก",
		"7": "เจ็ด",
		"8": "แปด",
		"9": "เก้า",
		"0": "",
	}
	unit := len(stringAmountMoney)
	var thaiNumber string
	for _, runeValue := range stringAmountMoney {
		if string(runeValue) != "0" {
			thaiNumber += fmt.Sprintf("%s%s", thaiCharactor[string(runeValue)], unitMap[unit])
		}
		unit--
	}

	return thaiNumber + "บาทถ้วน"
}
