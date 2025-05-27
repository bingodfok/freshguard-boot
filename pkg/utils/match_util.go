package utils

import "regexp"

const (
	phonePattern = `^((\+86)|(86))?(1[3-9]\d{9})$`
)

func PhoneMatch(phone string) bool {
	reg := regexp.MustCompile(phonePattern)
	return reg.MatchString(phone)
}
