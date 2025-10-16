package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.TrimPrefix(phoneNumber, "+1 ")
	phoneNumber = strings.TrimPrefix(phoneNumber, "1")
	var res strings.Builder
	for _, v := range phoneNumber {
		if v >= '0' && v <= '9' {
			res.WriteRune(v)
		}
	}
	phone := res.String()
	if len(phone) != 10 {
		return "", errors.New("")
	}
	if phone[0] < '2' || phone[3] < '2' {
		return "", errors.New("")
	}
	return phone, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phone[:3], phone[3:6], phone[6:]), nil
}
