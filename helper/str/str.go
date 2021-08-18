package str

import "strconv"

func StringToInt(text string) (res int) {
	res, _ = strconv.Atoi(text)
	return res
}

func StringToFloat64(s string) (res float64) {
	res, _ = strconv.ParseFloat(s, 64)
	return res // 3.14159265
}
