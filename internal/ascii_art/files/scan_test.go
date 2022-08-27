package files

import (
	"testing"
)

func TestScan(t *testing.T) {
	arg := "standard"
	res := map[rune][]string{
		'0': {"        ", "  ___   ", " / _ \\  ", "| | | | ", "| |_| | ", " \\___/  ", "        ", "        "},
	}
	got, _ := ReadTemplate(arg)

	comp := true
	for i, v := range res['0'] {
		if v != got['0'][i] {
			comp = false
			break
		}
	}
	if !comp {
		t.Errorf("got %q, wanted %q", got['0'], res['0'])
	}
}
