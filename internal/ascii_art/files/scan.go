package files

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

func ReadTemplate(arg string) (map[rune][]string, error) {
	switch arg {
	case "-sh", "shadow", "shadow.txt":
		arg = "internal/ascii_art/templates/shadow.txt"
	case "-th", "thinkertoy", "thinkertoy.txt":
		arg = "internal/ascii_art/templates/thinkertoy.txt"
	case "-st", "standard", "standard.txt":
		arg = "internal/ascii_art/templates/standard.txt"
	default:
		return nil, errors.New("invalid input")
	}

	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	template := map[rune][]string{}
	readFile := bufio.NewReader(file)
	text := ""
	for i := 32; i < 127; i++ {
		s, _, err := readFile.ReadLine()
		if err != nil {
			return nil, err
		}
		text += string(s) + "\n"
		for j := 0; j < 8; j++ {
			line, _, err := readFile.ReadLine()
			if err != nil {
				return nil, err
			}
			text += string(line) + "\n"
			template[rune(i)] = append(template[rune(i)], string(line))
		}
	}
	if !validTemplate(text, arg) {
		// log.Fatal("invalid template: template may have been changed")
		return nil, err
	}

	return template, nil
}

func validTemplate(text, arg string) bool {
	hashs := map[string]string{
		"internal/ascii_art/templates/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
		"internal/ascii_art/templates/thinkertoy.txt": "a57beec43fde6751ba1d30495b092658a064452f321e221d08c3ac34a9dc1294",
		"internal/ascii_art/templates/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text))) == hashs[arg]
}

func ValidText(text string) (string, bool) {
	symbols := ""
	found := false
	for _, v := range text {
		if v > 127 {
			symbols += string(v)
			found = true
		}
	}

	if found {
		return symbols, false
	}
	return text, true
}
