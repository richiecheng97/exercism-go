package cipher

import (
	"strings"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

func NewCaesar() Cipher {
	return NewShift(3)
}

type shift struct {
	distance int
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift{
		distance: distance,
	}
}

func (c shift) Encode(input string) string {
	input = strings.ToLower(input)
	var res strings.Builder
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+rune(c.distance)+26)%26
			res.WriteRune(char)
		}
	}
	return res.String()
}

func (c shift) Decode(input string) string {
	c.distance = -c.distance
	return c.Encode(input)
}

type vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	if key == "" || (key != "" && strings.Count(key, "a") == len(key)) {
		return nil
	}
	for _, char := range key {
		if char < 'a' || char > 'z' {
			return nil
		}
	}
	return vigenere{
		key: key,
	}
}

func (v vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	var result strings.Builder
	keyIndex := 0

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shift := int(v.key[keyIndex%len(v.key)] - 'a')
			encodedChar := 'a' + (char-'a'+rune(shift))%26
			result.WriteRune(encodedChar)
			keyIndex++
		}
	}

	return result.String()
}

func (v vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	var result strings.Builder
	keyIndex := 0

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shift := int(v.key[keyIndex%len(v.key)] - 'a')
			decodedChar := 'a' + (char-'a'-rune(shift)+26)%26
			result.WriteRune(decodedChar)
			keyIndex++
		}
	}

	return result.String()
}
