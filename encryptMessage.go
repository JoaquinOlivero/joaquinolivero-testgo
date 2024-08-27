package main

import "strings"

type Params struct {
	key string
	msg string
}

func EncryptMessage(params Params) string {
	if len(strings.TrimSpace(params.msg)) == 0 {
		return ""
	}

	if len(strings.TrimSpace(params.key)) == 0 {
		params.key = "dcj"
	}

	var str strings.Builder
	for i := 0; i < len(params.msg); i++ {
		letter := string(params.msg[i])
		if isVowel(letter) {
			str.WriteString(params.key + letter)
		} else {
			str.WriteString(letter)
		}

	}

	return str.String()
}

func isVowel(v string) bool {
	switch v {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	default:
		return false
	}
}
