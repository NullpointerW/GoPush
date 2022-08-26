package utils

import (
	"strconv"
	"time"
)

const TimeParseLayout = "2006-01-02 15:04:05"

func GenerateId(origin uint64) string {
	return strconv.FormatUint(origin, 10) +
		":" +
		strconv.FormatInt(time.Now().UnixNano(), 10)
}