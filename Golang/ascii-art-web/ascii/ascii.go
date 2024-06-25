package ascii

import (
	"os"
	"strings"
)

func Ascii(sentence string, banner string) string {
	base := []byte{}
	var err error
	switch banner {
	case "thinkertoy":
		base, err = os.ReadFile("./ascii/banner/thinkertoy.txt")
		if err != nil {
			return "ERROR"
		}
	case "shadow":
		base, err = os.ReadFile("./ascii/banner/shadow.txt")
		if err != nil {
			return "ERROR"
		}
	case "standard":
		base, err = os.ReadFile("./ascii/banner/standard.txt")
		if err != nil {
			return "ERROR"
		}
	}
	str_base := strings.Split(string(base), string('\n'))
	str := strings.Split(FormatText(sentence), string('\n'))
	str_output := ""
	for i := 0; i < len(str); i++ {
		if str[i] == "" {
			str_output += string('\n')
			continue
		}
		for k := 1; k < 9; k++ {
			for l := 0; l < len(str[i]); l++ {
				str_output += str_base[(int(str[i][l]-32))*9+k]
			}
			str_output += string('\n')
		}
	}
	output_file, _ := os.OpenFile("text.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	_, _ = output_file.WriteString(str_output)
	return str_output
}

func FormatText(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 13 {
		} else {
			str += string(s[i])
		}
	}
	return str
}
