package change

import (
	"strings"
)

type Printer struct {
	line     string
	template map[rune][]string
}

func ToBig(text string, template map[rune][]string) string {
	res := ""
	text = strings.Replace(text, "\r\n", "\n", -1)
	t := strings.Split(text, "\n")
	var p Printer
	p.template = template
	var notunique bool
	for i, c := range t {
		p.line = c
		if c == "" {
			if !notunique && i == len(t)-1 {
				break
			}
			res += "\n"
			continue
		}
		notunique = true
		res += p.print()
	}
	return res
}

func (p *Printer) print() string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, c := range p.line {
			res += string(p.template[c][i])
		}
		res += "\n"
	}
	return res
}
