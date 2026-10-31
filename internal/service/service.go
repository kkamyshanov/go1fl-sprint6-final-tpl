package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("файл не содержит текста")
	}
	ordinaryText := strings.ContainsFunc(input, func(r rune) bool {
		return r != '.' && r != '-' && !unicode.IsSpace(r)
	})
	if !ordinaryText {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
