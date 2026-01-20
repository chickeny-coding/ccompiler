package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State byte

const (
	_  State = iota
	NN       // NoNe
	RI       // Reading Identity
	RN       // Reading Number
)

type Infoer interface {
	info() string
	get() string
}

func (t Special) info() string {
	return "S" + tki[t]
}
func (t Special) get() string {
	return "S" + tki[t]
}
func (s Identity) info() string {
	return "I" + string(s)
}
func (s Identity) get() string {
	return "I" + string(s)
}
func (n Number) info() string {
	return "N" + strconv.Itoa(int(n))
}
func (n Number) get() string {
	return "N" + strconv.Itoa(int(n))
}

type Special byte
type Identity string
type Number int

const (
	_   Special = iota
	LPT         // ( Left ParenThesis
	RPT         // ) Right ParenThesis
	LSB         // [ Left Square Bracket
	RSB         // ] Left Square Bracket
	LCB         // { Left Curly Bracket
	RCB         // } Right Curly Bracket
)

var stt map[byte]Special = map[byte]Special{
	'(': LPT,
	')': RPT,
	'[': LSB,
	']': RSB,
	'{': LCB,
	'}': RCB,
}
var tks []Infoer = make([]Infoer, 0)
var tki [7]string = [7]string{
	"",
	"(",
	")",
	"[",
	"]",
	"{",
	"}",
}

func isSpecial(c byte) bool {
	return c == '(' || c == ')' || c == '[' || c == ']' || c == '{' || c == '}'
}

func isWhiteSpace(c byte) bool {
	return c <= ' '
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func lexer(fn string) ([]Infoer, error) {
	clt, err := os.Open(fn)
	if err != nil {
		return nil, fmt.Errorf("Failed opening file")
	}
	defer clt.Close()
	rd := bufio.NewReader(clt)
	c := byte(0)
	s := NN
	n := Number(0)
	str := Identity("")
	for {
		c, err = rd.ReadByte()
		if err != nil {
			break
		}
		switch s {
		case NN:
			if isSpecial(c) {
				tks = append(tks, stt[c])
			} else if isDigit(c) {
				n = Number(c) - '0'
				s = RN
			} else if !isWhiteSpace(c) {
				str = Identity(c)
				s = RI
			}
		case RI:
			if isSpecial(c) || isWhiteSpace(c) {
				tks = append(tks, str)
				if isSpecial(c) {
					tks = append(tks, stt[c])
				}
				s = NN
			} else {
				str = str + Identity(c)
			}
		case RN:
			if isDigit(c) {
				n = n*10 + Number(c) - '0'
			} else if isSpecial(c) || isWhiteSpace(c) {
				tks = append(tks, n)
				if isSpecial(c) {
					tks = append(tks, stt[c])
				}
				s = NN
			} else {
				return nil, fmt.Errorf("Identity starts with numbers")
			}
		default:
		}
	}
	return tks, nil
}
