package utils

import "time"

func GetStringFromPointer(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func PointerString(p string) *string {
	return &p
}

func PtoTime(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}