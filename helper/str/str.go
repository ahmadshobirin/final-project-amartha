package str

import "strconv"

func StringToInt(text string) (res int) {
	res, _ = strconv.Atoi(text)
	return res
}
