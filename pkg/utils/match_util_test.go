package utils

import (
	"testing"
)

func TestPhoneMatch(t *testing.T) {
	t.Log(PhoneMatch("136993609826"))
	t.Log(PhoneMatch("236983602826"))
	t.Log(PhoneMatch("13698360226"))
	t.Log(PhoneMatch("15520568045"))
}
