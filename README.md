# ansistrings

# Description

ansistrings provide 3 types

* ANSIStyle
  * define ANSI escaping for string
* ANSIString
  * string applyed ANSIStyle definition
* ANSIStrings
  * has set of ANSIString

ANSIStrings Str method add new ANSIString to ANSIStrings.
You can use all ANSIStyle methods from ANSIStrings.
Their methods are applyed to last creating ANSIString.
And then ANSIStrings Print method print ANSI escaped strings 
and discard them from ANSIStrings.

```
 v := NewANSIStrings()
 // add new ANSIString
 v.Str("test")
 // Bold() affects "test" 
 v.Bold()
 // add another ANSIString
 v.Str("test2")
 // Blue() affects "test2" and "\033[0m" will be added after "test"
 // so, Bold() won't affect "test2"
 v.Blue()
 // print all of the above. v has no ANSIString.
 v.Print()

 // If you don't want to discard created ANSIString, use String() instead.
 fmt.Print(v.String())
 // is equal to
 fmt.Print(v)
```

# Example
```
	import (
        "fmt"
        s "github.com/ktat/go-ansistrings"
    )

    str := NewANSIString("TEST")
    str.Bold().Blue()
    // print bold and blue "TEST"
    fmt.Println(str)

    v := NewANSIStrings()
	step := 10
	for r := 0; r < 256; r += step {
		for g := 0; g < 256; g += step {
			for b := 0; b < 256; b += step {
                // set RGB colored string
				v.Str("*").RGB(r, g, b)
			}
		}
	}
    // print them
    v.Print()

    // define styles to reuse
  	style := s.NewANSIStyle()
	style.Bold().UnderLine().Blue()
	style2 := s.NewANSIStyle()
	style2.Faint().Delete().RGB(200, 100, 50).BgRGB(50, 100, 200)

    // clear console
	v.Clear().
		Str("1: Bold string and Down after this output").
		Bold().
		Down().
		Pause(100).
		Back(20).
		Str("2: Back 20").
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
		Str("13: Style2\n Test 2\n").Style(style2).
		Print()
```

# Reference

https://en.wikipedia.org/wiki/ANSI_escape_code

# Author

Atsushi Kato (ktat)

# License

MIT: https://ktat.mit-license.org/2016

