package main

import (
	"fmt"
	"time"

	s "github.com/ktat/go-ansistrings"
)

func main() {
	a := s.NewANSIString("TEST")
	a.Bold().White().UnderLine().Delete().BgColor(s.Red)
	fmt.Println(a)

	v := s.NewANSIStrings()

	step := 10
	for r := 0; r < 256; r += step {
		for g := 0; g < 256; g += step {
			for b := 0; b < 256; b += step {
				v.Str("*").RGB(r, g, b)
			}
		}
	}
	v.Print().Pause(100)
	for r := 0; r < 256; r += step {
		for g := 0; g < 256; g += step {
			for b := 0; b < 256; b += step {
				v.Str(" ").BgRGB(r, g, b)
			}
		}
	}
	v.Print().Pause(100).Clear().Print()
	for i := 1; i < 256; i++ {
		v.Str(fmt.Sprintf("(%03d)", i)).
			ColorN(i).
			Print()
	}
	for i := 1; i < 256; i++ {
		v.Str(fmt.Sprintf("(%03d)", i)).
			BgColorN(i).
			Print()
	}
	v.Pause(100).Clear().Print()

	h := 10
	w := 24
	c := 0
	for i := 1; i < w; i++ {
		for j := 1; j < h; j++ {
			v.Print()
			v.Str(fmt.Sprintf("%03d", c))
			if c%10 == 0 {
				v.Bold()
			} else if c%10 == 1 {
				v.Faint()
			} else if c%10 == 2 {
				v.Italic()
			} else if c%10 == 3 {
				v.UnderLine()
			} else if c%10 == 4 {
				v.Blink()
			} else if c%10 == 5 {
				v.RapidBlink()
			} else if c%10 == 6 {
				v.Inverted()
			} else if c%10 == 7 {
				v.Conceal()
			} else if c%10 == 8 {
				v.Delete()
			} else if c%10 == 9 {
				v.Bold()
			}
			v.Font(c%14 + 1).
				ColorN(c).
				Down(1).
				Back(3).
				Pause(5)
			c++
		}
		v.Forward(3)
		v.Up(h)
		v.Print()
	}
	v.Pos(1, 11).Str("End\n").Pause(2000).Print()

	style := s.NewANSIStyle()
	style.Bold().UnderLine().Blue()

	style2 := s.NewANSIStyle()
	style2.Faint().Delete().RGB(200, 100, 50).BgRGB(50, 100, 200)
	fmt.Println("AAA")
	v.Clear().
		Str("1: Bold string and Down after this output").
		Bold().
		Down().
		Pause(1000).
		Back(20).
		Str("2: back 20").
		Down().
		Str("3: Blue string and Down 10 after this output").
		Blue().
		Down(10).
		Str("4: Downed 10").
		Pos(5, 9).
		Str("5: Pos 5, 9").
		Pos(6, 10).
		Str("6: Pos 6, 10").
		Pos(7, 11).
		Str("7: Pos 7, 11").
		Pos(6, 12).
		Str("8: Pos 6, 12").
		Pos(1, 13).
		Str("9: Pos 1,13\n").
		Str("10: Style1\n Test 1\n").Style(style).
		Str("11: Style2\n Test 1\n").Style(style2).
		Str("12: Style1\n Test 2\n").Style(style).
		Str("13: Style2\n Test 2\n").Style(style2)
	fmt.Println("BBB")
	fmt.Printf("%#v", v.String())
	time.Sleep(3 * time.Second)
	v.Print()
}
