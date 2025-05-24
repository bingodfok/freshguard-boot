package test

import (
	"encoding/json"
	"fmt"
	"freshguard-boot/pkg/auth"
	"testing"
)

func TestGenToken(t *testing.T) {
	jwtAuth := auth.JwtAuth{
		SigningKey: "AllYourBase",
	}
	claims := auth.StandardClaims{
		UserId:   "sasa",
		UserName: "sasa",
	}
	token, err := jwtAuth.GenToken(claims)
	if err != nil {
		t.Errorf("Error generating token: %v", err)
	}
	fmt.Println(token)
	c, err := jwtAuth.ParseToken(token)
	if err != nil {
		t.Errorf("Error parsing token: %v", err)
	}
	marshal, err := json.Marshal(c)
	if err != nil {
		t.Errorf("Error marshalling token: %v", err)
	}
	fmt.Println(string(marshal))

}
