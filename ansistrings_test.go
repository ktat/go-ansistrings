package ansistrings_test

import (
	"testing"

	s "github.com/ktat/go-ansistrings"
)

var v = s.NewANSIStrings()

func init() {
	v.Str("test")
}

func TestColorNumFromName(t *testing.T) {
	n := 30
	testNames := []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "light_gray"}
	for _, c := range testNames {
		cn, e := s.ColorNumFromName(c)
		if cn != n {
			t.Errorf("Get %d, want %d: %#v", cn, n, e)
		}
		n++
	}
	testNames = []string{"dark_gray", "light_red", "light_green", "light_yellow", "light_blue", "light_magenta", "light_cyan", "white"}
	n = 90
	for _, c := range testNames {
		cn, e := s.ColorNumFromName(c)
		if cn != n {
			t.Errorf("Get %d, want %d: %#v", cn, n, e)
		}
		n++
	}
}

func TestBgColor(t *testing.T) {
	v.BgColor(s.Black)
	a := "\033[40mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}

	a = "\033[107mtest\033[0m"
	v.BgColor(s.White)
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}

	v.CurrentStr().UnsetBgColor()
	if v.String() != "test" {
		t.Errorf("Get %#v, want test", v.String())
	}
}

func TestBgColorN(t *testing.T) {
	v.BgColorN(125)
	a := "\033[48;5;125mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
	v.CurrentStr().UnsetBgColor()
	if v.String() != "test" {
		t.Errorf("Get %#v, want test", v.String())
	}
}

func TestBlack(t *testing.T) {
	v.Black()
	a := "\033[30mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestRed(t *testing.T) {
	v.Red()
	a := "\033[31mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestColor(t *testing.T) {
	v.ColorN(125)
	a := "\033[38;5;125mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestColor2(t *testing.T) {
	v.ColorN(30)
	a := "\033[38;5;30mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestColor3(t *testing.T) {
	v.ColorN(90)
	a := "\033[38;5;90mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestBold(t *testing.T) {
	v.CurrentStr().UnsetColor()
	if v.String() != "test" {
		t.Errorf("Get %#v, want test", v.String())
	}

	v.Bold()
	a := "\033[1mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestUnderLine(t *testing.T) {
	v.ResetStyle()
	v.UnderLine()
	a := "\033[4mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestBlink(t *testing.T) {
	v.ResetStyle()
	v.Blink()
	a := "\033[5mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestInverted(t *testing.T) {
	v.ResetStyle()
	v.Inverted()
	a := "\033[7mtest\033[0m"
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestAll(t *testing.T) {
	v.ResetStyle()
	a := "\033[31m\033[42m\033[1m\033[4m\033[5mtest\033[0m"

	v.Red()
	v.UnderLine()
	v.Blink()
	v.BgColor(s.Green)
	v.Bold()
	if v.String() != a {
		t.Errorf("Get %#v, want %#v", v.String(), a)
	}
}

func TestReset(t *testing.T) {
	v.ResetStyle()
	var v2 = s.ANSIString{Str: "test"}
	if v.RawString() != v2.RawString() {
		t.Errorf("Get %#v, want %#v", v, v2)
	}
}

func TestChain(t *testing.T) {
	v.ResetStyle()
	v := v.Blue().Str("test2").Bold().Str("test3").Cyan().String()
	v2 := "\033[34mtest\033[0m\033[1mtest2\033[0m\033[36mtest3\033[0m"
	if v != v2 {
		t.Errorf("Get %#v, want %#v", v, v2)
	}
}
