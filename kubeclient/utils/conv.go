package utils

import (
	"fmt"
	"time"
	"strconv"
)

func GetLocalTime() string {
	t := time.Now().Local()
	return fmt.Sprintf("%v", t.Format("2006-01-02 15:04:05 -0700"))
}

func GetDate() string {
	t := time.Now().Local()
	return fmt.Sprintf("%v", t.Format("2006-01-02"))
}

func GetTimestrByIndex(idx int) string {
	if idx < 0 || idx >= 1440 {
		return ""
	}
	m := idx%60
	h := idx/60
	return fmt.Sprintf("%02v:%02v", h, m)
}

func GetIndexedTime(t time.Time) int {
	h, _ := strconv.Atoi(fmt.Sprintf("%v", t.Format("15")))
	m, _ := strconv.Atoi(fmt.Sprintf("%v", t.Format("04")))
	retv := h*60
	retv += m
	return retv
}