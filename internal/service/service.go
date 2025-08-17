package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(data string) bool {
	validChars := ".- "
	for _, ch := range data {
		if !strings.ContainsRune(validChars, ch) {
			return false
		}
	}
	return true
}

func Convert(data string) string {
	if isMorse(data) {
		return morse.ToText(data)
	}
	return morse.ToText(data)
}
