package client

import (
	"strconv"
)

func StrInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}
