// Package ansistrings returns ANSI escaped string
package ansistrings

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// constant value of colors
const (
	Black        = 30
	Red          = 31
	Green        = 32
	Yellow       = 33
	Blue         = 34
	Magenta      = 35
	Cyan         = 36
	LightGray    = 37
	DarkGray     = 90
	LightRed     = 91
	LightGreen   = 92
	LightYellow  = 93
	LightBlue    = 94
	LightMagenta = 95
	LightCyan    = 96
	White        = 97
	_bold        = "\033[1m"
	_faint       = "\033[2m"
	_italic      = "\033[3m"
	_underLine   = "\033[4m"
	_blink       = "\033[5m"
	_rapidBlink  = "\033[6m"
	_inverted    = "\033[7m"
	_conceal     = "\033[8m"
	_delete      = "\033[9m"
	_reset       = "\033[0m"
	_up          = "A"
	_down        = "B"
	_forward     = "C"
	_back        = "D"
)

var name2color = map[string]int{
	"black":         Black,
	"red":           Red,
	"green":         Green,
	"yellow":        Yellow,
	"blue":          Blue,
	"magenta":       Magenta,
	"cyan":          Cyan,
	"light_gray":    LightGray,
	"dark_gray":     DarkGray,
	"light_red":     LightRed,
	"light_green":   LightGreen,
	"light_yellow":  LightYellow,
	"light_blue":    LightBlue,
	"light_magenta": LightMagenta,
	"light_cyan":    LightCyan,
	"white":         White,
}

// ANSIStyle is set of ANSI escape settings
type ANSIStyle struct {
	color struct {
		color int
		isSet bool
	}
	colorN struct {
		color int
		isSet bool
	}
	bgColor struct {
		color int
		isSet bool
	}
	rgb struct {
		r     int
		g     int
		b     int
		isSet bool
	}
	bgRgb struct {
		r     int
		g     int
		b     int
		isSet bool
	}
	bgColorN struct {
		color int
		isSet bool
	}
	withBold       bool
	withDelete     bool
	withItalic     bool
	withBlink      bool
	withRapidBlink bool
	withFaint      bool
	withConceal    bool
	withInverted   bool
	withUnderLine  bool
	font           int
	sleep          time.Duration
	skipSleep      bool
}

// ANSIString is struct which contains string with ANSI escaping setting
type ANSIString struct {
	Str       string
	direction struct {
		n         int
		direction string
	}
	position struct {
		x int
		y int
	}
	clear bool
	ANSIStyle
}

// ANSIStrings is struct which contains ANSIString
type ANSIStrings struct {
	strings []ANSIString
	index   int
}

// NewANSIStrings returns new ANSIStrings
func NewANSIStrings() ANSIStrings {
	return ANSIStrings{}
}

// NewANSIStyle returns new ANSIStyle
func NewANSIStyle() ANSIStyle {
	return ANSIStyle{}
}

// NewANSIString returns new ANSIString
func NewANSIString(s string) ANSIString {
	return ANSIString{Str: s}
}

// ColorNumFromName return number of colors. color name shoud be lower case.
func ColorNumFromName(color string) (cn int, e error) {
	v, ok := name2color[color]
	if ok {
		cn = v
	} else {
		e = errors.New("unknown color name: " + color)
	}
	return cn, e
}

// Print prints ANSI escaped strings and disard them
func (s *ANSIStrings) Print() *ANSIStrings {
	str := ""
	for _, as := range s.strings {
		str = as.String()
		if str != "" {
			fmt.Print(str)
		}
	}
	s.strings = make([]ANSIString, 0)
	return s
}

// Style sets string predefined ANSIStyle
func (s *ANSIStrings) Style(style ANSIStyle) *ANSIStrings {
	as := s.CurrentStr()
	as.color = style.color
	as.colorN = style.colorN
	as.rgb = style.rgb
	as.bgRgb = style.bgRgb
	as.bgColor = style.bgColor
	as.bgColorN = style.bgColorN
	as.withBold = style.withBold
	as.withDelete = style.withDelete
	as.withItalic = style.withItalic
	as.withBlink = style.withBlink
	as.withRapidBlink = style.withRapidBlink
	as.withFaint = style.withFaint
	as.withConceal = style.withConceal
	as.withInverted = style.withInverted
	as.withUnderLine = style.withUnderLine
	as.font = style.font
	as.sleep = style.sleep
	return s
}

// Pause sleeps given millisecond(s)
func (s *ANSIStrings) Pause(i ...time.Duration) *ANSIStrings {
	s.Str("")
	s.CurrentStr().Pause(i...)
	return s
}

// Pause sleeps given millisecond(s)
func (s *ANSIStyle) Pause(i ...time.Duration) *ANSIStyle {
	var n time.Duration = 1
	if len(i) != 0 && i[0] > 0 {
		n = i[0]
	}
	s.sleep = n * time.Millisecond
	return s
}

// RawString returns raw string
func (s *ANSIStrings) RawString() string {
	str := ""
	for i := 0; i < len(s.strings); i++ {
		str += fmt.Sprintf("%#v", s.strings[i].String())
	}
	return str
}

// RawString retruns raw string
func (s *ANSIString) RawString() string {
	return fmt.Sprintf("%#v", s.String())
}

// String returns ANSI escaped string.
func (s ANSIStrings) String() string {
	str := ""
	for i := 0; i < len(s.strings); i++ {
		as := s.strings[i]
		as.skipSleep = true
		str += as.String()
		as.skipSleep = false
	}
	return str
}

// String returns ANSI escaped string
func (s ANSIString) String() string {
	color := ""
	if s.clear {
		color += "\033[2J\033[1;1H"
	} else if s.position.x != 0 {
		color += fmt.Sprintf("\033[%d;%dH", s.position.y, s.position.x)
	} else if s.direction.n != 0 {
		color += fmt.Sprintf("\033[%d%s", s.direction.n, s.direction.direction)
	} else if s.sleep != 0 && s.skipSleep == false {
		time.Sleep(s.sleep)
	} else {
		if s.color.isSet {
			color += fmt.Sprintf("\033[%dm", s.color.color)
		} else if s.rgb.isSet {
			color += fmt.Sprintf("\033[38;2;%d;%d;%dm", s.rgb.r, s.rgb.g, s.rgb.b)
		} else if s.colorN.isSet {
			color += fmt.Sprintf("\033[38;5;%dm", s.colorN.color)
		}
		if s.bgColor.isSet {
			color += fmt.Sprintf("\033[%dm", s.bgColor.color+10)
		} else if s.bgRgb.isSet {
			color += fmt.Sprintf("\033[48;2;%d;%d;%dm", s.bgRgb.r, s.bgRgb.g, s.bgRgb.b)
		} else if s.bgColorN.isSet {
			color += fmt.Sprintf("\033[48;5;%dm", s.bgColorN.color)
		}
		if s.withBold {
			color += _bold
		}
		if s.withFaint {
			color += _faint
		}
		if s.withItalic {
			color += _italic
		}
		if s.withUnderLine {
			color += _underLine
		}
		if s.withBlink {
			color += _blink
		}
		if s.withRapidBlink {
			color += _rapidBlink
		}
		if s.withInverted {
			color += _inverted
		}
		if s.withConceal {
			color += _conceal
		}
		if s.withDelete {
			color += _delete
		}
		if s.font != 0 {
			color += fmt.Sprintf("\033[%dm", s.font+10)
		}

		if color != "" {
			return fmt.Sprintf(color+"%s"+_reset, s.resetWithLineBreak(s.Str, color))
		}
	}
	// not required reset with position
	if color != "" {
		return fmt.Sprintf(color+"%s", s.Str)
	}
	return s.Str
}

func (s *ANSIString) resetWithLineBreak(str string, color string) string {
	var r = regexp.MustCompile("(?s)([\r\n]+)")
	return r.ReplaceAllString(str, _reset+"$1"+color)
}

// BgColor sets background color of string
func (s *ANSIStrings) BgColor(c int) *ANSIStrings {
	s.CurrentStr().BgColor(c)
	return s
}

// BgColor sets background color of string
func (s *ANSIStyle) BgColor(c int) *ANSIStyle {
	s.bgColor.color = c
	s.bgColor.isSet = true
	s.bgColorN.isSet = false
	s.bgRgb.isSet = false
	return s
}

// BgColorN sets background color of string
func (s *ANSIStrings) BgColorN(c int) *ANSIStrings {
	if c < 0 || c > 255 {
		panic("invalid argument: valid range is from 0 to 255")
	}
	s.CurrentStr().BgColorN(c)
	return s
}

// BgColorN sets background color of string
func (s *ANSIStyle) BgColorN(c int) *ANSIStyle {
	if c < 0 || c > 255 {
		panic("invalid argument: valid range is from 0 to 255")
	}
	s.bgColorN.color = c
	s.bgColor.isSet = false
	s.bgColorN.isSet = true
	s.bgRgb.isSet = false
	return s
}

// UnsetColor unsets color of string
func (s *ANSIStyle) UnsetColor() *ANSIStyle {
	s.color.isSet = false
	s.colorN.isSet = false
	s.rgb.isSet = false
	return s
}

// UnsetBgColor unsets background color of string
func (s *ANSIStyle) UnsetBgColor() *ANSIStyle {
	s.bgColor.isSet = false
	s.bgColorN.isSet = false
	s.bgRgb.isSet = false
	return s
}

// Color set string color
func (s *ANSIStyle) Color(color int) *ANSIStyle {
	s.colorN.isSet = false
	s.rgb.isSet = false
	s.color.isSet = true
	s.color.color = color
	return s
}

// Color set color
func (s *ANSIStrings) Color(color int) *ANSIStrings {
	s.CurrentStr().Color(color)
	return s
}

// Black sets string black
func (s *ANSIStrings) Black() *ANSIStrings {
	s.CurrentStr().Black()
	return s
}

// Black sets string black
func (s *ANSIStyle) Black() *ANSIStyle {
	s.Color(Black)
	return s
}

// Red sets string red
func (s *ANSIStrings) Red() *ANSIStrings {
	s.CurrentStr().Red()
	return s
}

// Red sets string red
func (s *ANSIStyle) Red() *ANSIStyle {
	s.Color(Red)
	return s
}

// Green set string green
func (s *ANSIStrings) Green() *ANSIStrings {
	s.CurrentStr().Green()
	return s
}

// Green sets string green
func (s *ANSIStyle) Green() *ANSIStyle {
	s.Color(Green)
	return s
}

// Yellow sets string yellow
func (s *ANSIStrings) Yellow() *ANSIStrings {
	s.CurrentStr().Yellow()
	return s
}

// Yellow sets string yellow
func (s *ANSIStyle) Yellow() *ANSIStyle {
	s.Color(Yellow)
	return s
}

// Blue sets string blue
func (s *ANSIStrings) Blue() *ANSIStrings {
	s.CurrentStr().Blue()
	return s
}

// Blue sets string blue
func (s *ANSIStyle) Blue() *ANSIStyle {
	s.Color(Blue)
	return s
}

// Magenta sets string magenta
func (s *ANSIStrings) Magenta() *ANSIStrings {
	s.CurrentStr().Magenta()
	return s
}

// Magenta sets string magenta
func (s *ANSIStyle) Magenta() *ANSIStyle {
	s.Color(Magenta)
	return s
}

// Cyan sets string cyan
func (s *ANSIStrings) Cyan() *ANSIStrings {
	s.CurrentStr().Cyan()
	return s
}

// Cyan sets string cyan
func (s *ANSIStyle) Cyan() *ANSIStyle {
	s.Color(Cyan)
	return s
}

// White sets string white
func (s *ANSIStrings) White() *ANSIStrings {
	s.CurrentStr().White()
	return s
}

// White sets string white
func (s *ANSIStyle) White() *ANSIStyle {
	s.Color(White)
	return s
}

// LightRed sets string light red
func (s *ANSIStrings) LightRed() *ANSIStrings {
	s.CurrentStr().LightRed()
	return s
}

// LightRed sets string light red
func (s *ANSIStyle) LightRed() *ANSIStyle {
	s.Color(LightRed)
	return s
}

// LightGreen sets string light green
func (s *ANSIStrings) LightGreen() *ANSIStrings {
	s.CurrentStr().LightGreen()
	return s
}

// LightGreen sets string light green
func (s *ANSIStyle) LightGreen() *ANSIStyle {
	s.Color(LightGreen)
	return s
}

// LightYellow sets string light yellow
func (s *ANSIStrings) LightYellow() *ANSIStrings {
	s.CurrentStr().LightYellow()
	return s
}

// LightYellow sets string light yellow
func (s *ANSIStyle) LightYellow() *ANSIStyle {
	s.Color(LightYellow)
	return s
}

// LightBlue sets string light blue
func (s *ANSIStrings) LightBlue() *ANSIStrings {
	s.CurrentStr().LightBlue()
	return s
}

// LightBlue sets string light blue
func (s *ANSIStyle) LightBlue() *ANSIStyle {
	s.Color(LightBlue)
	return s
}

// LightMagenta sets string light magenta
func (s *ANSIStrings) LightMagenta() *ANSIStrings {
	s.CurrentStr().LightMagenta()
	return s
}

// LightMagenta sets string light magenta
func (s *ANSIStyle) LightMagenta() *ANSIStyle {
	s.Color(LightMagenta)
	return s
}

// LightCyan sets string light cyan
func (s *ANSIStrings) LightCyan() *ANSIStrings {
	s.CurrentStr().LightCyan()
	return s
}

// LightCyan sets string light cyan
func (s *ANSIStyle) LightCyan() *ANSIStyle {
	s.Color(LightCyan)
	return s
}

// LightGray sets string light gray
func (s *ANSIStrings) LightGray() *ANSIStrings {
	s.CurrentStr().LightGray()
	return s
}

// LightGray sets string light gray
func (s *ANSIStyle) LightGray() *ANSIStyle {
	s.Color(LightGray)
	return s
}

// DarkGray sets string dark gray
func (s *ANSIStrings) DarkGray() *ANSIStrings {
	s.CurrentStr().DarkGray()
	return s
}

// DarkGray sets string dark gray
func (s *ANSIStyle) DarkGray() *ANSIStyle {
	s.Color(DarkGray)
	return s
}

// ColorN sets string given color(argument range is from 0 to 255)
func (s *ANSIStrings) ColorN(c int) *ANSIStrings {
	s.CurrentStr().ColorN(c)
	return s
}

// ColorN sets string given color(argument range is from 0 to 255)
func (s *ANSIStyle) ColorN(c int) *ANSIStyle {
	if c < 0 || c > 255 {
		panic("invalid argument: valid range is from 0 to 255")
	}
	s.rgb.isSet = false
	s.color.isSet = false
	s.colorN.color = c
	s.colorN.isSet = true
	return s
}

// RGB set string given RGB color(argument range is from 0 to 255)
func (s *ANSIStrings) RGB(r int, g int, b int) *ANSIStrings {
	s.CurrentStr().RGB(r, g, b)
	return s
}

// RGB sets string given RGB color(argument range is from 0 to 255)
func (s *ANSIStyle) RGB(r int, g int, b int) *ANSIStyle {
	if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
		panic("invalid argument: valid range is from 0 to 255")
	}
	s.color.isSet = false
	s.colorN.isSet = false
	s.rgb.isSet = true
	s.rgb.r = r
	s.rgb.g = g
	s.rgb.b = b
	return s
}

// BgRGB sets string given RGB color(argument range is from 0 to 255)
func (s *ANSIStrings) BgRGB(r int, g int, b int) *ANSIStrings {
	s.CurrentStr().BgRGB(r, g, b)
	return s
}

// BgRGB sets string given RGB color(argument range is from 0 to 255)
func (s *ANSIStyle) BgRGB(r int, g int, b int) *ANSIStyle {
	if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
		panic("invalid argument: valid range is from 0 to 255")
	}
	s.bgColor.isSet = false
	s.bgColorN.isSet = false
	s.bgRgb.isSet = true
	s.bgRgb.r = r
	s.bgRgb.g = g
	s.bgRgb.b = b
	return s
}

// Bold sets string bold
func (s *ANSIStrings) Bold() *ANSIStrings {
	s.CurrentStr().Bold()
	return s
}

// Bold sets string bold
func (s *ANSIStyle) Bold() *ANSIStyle {
	s.withBold = true
	return s
}

// Faint sets string faint
func (s *ANSIStrings) Faint() *ANSIStrings {
	s.CurrentStr().Faint()
	return s
}

// Faint sets string faint
func (s *ANSIStyle) Faint() *ANSIStyle {
	s.withBold = true
	return s
}

// Italic sets string italic
func (s *ANSIStrings) Italic() *ANSIStrings {
	s.CurrentStr().Italic()
	return s
}

// Italic sets string italic
func (s *ANSIStyle) Italic() *ANSIStyle {
	s.withItalic = true
	return s
}

// Blink sets string blink
func (s *ANSIStrings) Blink() *ANSIStrings {
	s.CurrentStr().Blink()
	return s
}

// Blink sets string blink
func (s *ANSIStyle) Blink() *ANSIStyle {
	s.withBlink = true
	return s
}

// RapidBlink sets string rapid blink
func (s *ANSIStrings) RapidBlink() *ANSIStrings {
	s.CurrentStr().RapidBlink()
	return s
}

// RapidBlink sets string rapid blink
func (s *ANSIStyle) RapidBlink() *ANSIStyle {
	s.withRapidBlink = true
	return s
}

// Inverted sets string inverted
func (s *ANSIStrings) Inverted() *ANSIStrings {
	s.CurrentStr().Inverted()
	return s
}

// Inverted sets string inverted
func (s *ANSIStyle) Inverted() *ANSIStyle {
	s.withInverted = true
	return s
}

// UnderLine sets string underline
func (s *ANSIStrings) UnderLine() *ANSIStrings {
	s.CurrentStr().UnderLine()
	return s
}

// UnderLine sets string underline
func (s *ANSIStyle) UnderLine() *ANSIStyle {
	s.withUnderLine = true
	return s
}

// Delete sets string crossed out
func (s *ANSIStrings) Delete() *ANSIStrings {
	s.CurrentStr().Delete()
	return s
}

// Delete sets string crossed out
func (s *ANSIStyle) Delete() *ANSIStyle {
	s.withDelete = true
	return s
}

// Conceal sets string concealed
func (s *ANSIStrings) Conceal(t ...bool) *ANSIStrings {
	s.CurrentStr().Conceal()

	return s
}

// Conceal sets string concealed
func (s *ANSIStyle) Conceal() *ANSIStyle {
	s.withConceal = true
	return s
}

// Up curosr
func (s *ANSIStrings) Up(n ...int) *ANSIStrings {
	return s.setDirection(_up, n)
}

// Down curosr
func (s *ANSIStrings) Down(n ...int) *ANSIStrings {
	return s.setDirection(_down, n)
}

// Forward curosr
func (s *ANSIStrings) Forward(n ...int) *ANSIStrings {
	return s.setDirection(_forward, n)
}

// Back curosr
func (s *ANSIStrings) Back(n ...int) *ANSIStrings {
	return s.setDirection(_back, n)
}

func (s *ANSIStrings) setDirection(d string, n []int) *ANSIStrings {
	s.Str("")
	s.CurrentStr().direction.direction = d
	if len(n) == 0 || n[0] == 0 {
		s.CurrentStr().direction.n = 1
	} else {
		s.CurrentStr().direction.n = n[0]
	}
	return s
}

// Font changes font(from 1 to 14)
func (s *ANSIStrings) Font(n int) *ANSIStrings {
	if n < 1 || n > 14 {
		panic("n should be from 1 to 14")
	}
	s.CurrentStr().font = n
	return s
}

// Pos sets cursor given position.
func (s *ANSIStrings) Pos(x int, y int) *ANSIStrings {
	s.Str("")
	if x < 1 || y < 1 {
		panic("row and column should be greater than 0")
	}
	s.CurrentStr().position.x = x
	s.CurrentStr().position.y = y
	return s
}

// Clear clears console
func (s *ANSIStrings) Clear() *ANSIStrings {
	s.Str("")
	s.CurrentStr().clear = true
	return s
}

// ResetStyle resets ANSI escaping
func (s *ANSIString) ResetStyle() {
	str := s.Str
	*s = ANSIString{Str: str}
}

// ResetStyle resets ANSI escaping
func (s *ANSIStrings) ResetStyle() *ANSIStrings {
	str := s.CurrentStr().Str
	s.strings[s.index] = ANSIString{Str: str}
	return s
}

// ResetAll resets All strings' ANSI escaping
func (s *ANSIStrings) ResetAll() *ANSIStrings {
	for i := 0; i < len(s.strings); i++ {
		s.CurrentStr().ResetStyle()
	}
	return s
}

// Str add new ANSIString
func (s *ANSIStrings) Str(str string) *ANSIStrings {
	s.strings = append(s.strings, ANSIString{Str: str})
	s.index = len(s.strings) - 1
	return s
}

// CurrentStr return last set ANSIString
func (s *ANSIStrings) CurrentStr() *ANSIString {
	if len(s.strings) == 0 {
		panic("use Str() at first")
	}
	return &s.strings[s.index]
}
