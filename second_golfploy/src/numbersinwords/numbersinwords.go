package numbersinwords

import "strconv"

func ConvertNumberToThaiCharactor(inputNumber float64) string {
	var output string
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
	unitThai := map[int]string{
		6: "แสน",
		5: "หมื่น",
		4: "พัน",
		3: "ร้อย",
		2: "สิบ",
		1: "",
	}
	number := int(inputNumber)
	stringNumber := strconv.Itoa(number)
	unit := len(stringNumber)
	for _, numberCharactor := range stringNumber {
		if string(numberCharactor) != "0" {
			output += numberThai[string(numberCharactor)] + unitThai[unit]
		}
		unit--
	}
	return output + "บาทถ้วน"
}
