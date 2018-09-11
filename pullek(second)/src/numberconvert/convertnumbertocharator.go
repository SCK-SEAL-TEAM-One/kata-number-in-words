package numberconvert

import (
	"fmt"
)

func NumberConvert(number int) string {
	numbrtString := fmt.Sprintf("%d", number)
	unitNumber := map[int]string{
		2: "สิบ",
		3: "ร้อย",
		4: "พัน",
		5: "หมื่น",
		6: "แสน",
		7: "ล้าน",
	}
	numberThai := map[string]string{
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
	var thaiCharactor string
	lengthNumber := len(numbrtString)
	for _, units := range numbrtString {
		if string(units) != "0" {
			thaiCharactor += fmt.Sprintf("%s%s", numberThai[string(units)], unitNumber[lengthNumber])
		}
		lengthNumber--
	}
	return thaiCharactor + "บาทถ้วน"
}
