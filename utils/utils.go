package utils

import "strconv"

func StringToUint32(str string) (result uint32, err error) {
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {

		return
	}

	return uint32(u64), nil
}
