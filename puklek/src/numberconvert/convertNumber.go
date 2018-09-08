package numberconvert

import (
	"fmt"
)

func ConvertNumber(money float64) string {
	unitNumber := map[int]string{
		1: "",
		2: "สิบ",
		3: "ร้อย",
		4: "พัน",
		5: "หมื่น",
		6: "แสน",
	}
	number := map[string]string{
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

	var convertNumber string

	moneyString := fmt.Sprintf("%d", int(money))

	units := len(moneyString)

	for _, unitMoney := range moneyString {
		if string(unitMoney) != "0" {
			convertNumber += fmt.Sprintf("%s%s", number[string(unitMoney)], unitNumber[units])
			units--
		}
	}

	return convertNumber + "บาทถ้วน"
}
