package ascii_art

import (
	"errors"
	"fmt"

	"a-art-w/internal/ascii_art/files"

	"a-art-w/internal/ascii_art/change"
)

func AsciiArt(text, font []string) (string, error) {
	symbols, valid := files.ValidText(text[0])
	if !valid {
		fmt.Printf("found invalid character(s): %v\nplease change it\n", symbols)
		return "", errors.New("")
	}

	template, err := files.ReadTemplate(font[0])
	if err != nil {
		fmt.Println(fmt.Errorf("ReadTemplate: %v", err))
		return "", err
	}
	return change.ToBig(text[0], template), nil
}
