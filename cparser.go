package main

import (
	"fmt"
)

type Symbol byte

const (
	_ Symbol = iota
	PROGRAM
	BLOCK
	ELEMENT
	COMMENT
)

var ss [5]string = [5]string{
	"",
	"PROGRAM",
	"BLOCK",
	"ELEMENT",
	"COMMENT",
}

type SymbolGetter interface {
	get() string
}

func (s Symbol) get() string {
	return ss[s]
}

type Node struct {
	v SymbolGetter
	c []*Node
}

var lex []Infoer
var lex_i uint32
var itm map[byte]string = map[byte]string{
	'S': "Special",
	'I': "Identity",
	'N': "Number",
}

func move() {
	lex_i++
}

func match(c byte) (Infoer, error) {
	if lex[lex_i] == nil {
		[]int{}[0] = 0
		return nil, fmt.Errorf("Reached EOF")
	}
	if lex[lex_i].info()[0] != c {
		return nil, fmt.Errorf("Expected %s, got %s", itm[c], lex[lex_i].info()[1:])
	}
	v := lex[lex_i]
	move()
	return v, nil
}

func program(n *Node) error {
	n.v = PROGRAM
	if lex[lex_i] == nil {
		return nil
	}
	n.c = []*Node{{}, {}, {}, {}, {}, {}, {}}
	val, err := match('S')
	if err != nil {
		return err
	}
	if val.info()[1] != '(' {
		return fmt.Errorf("Expected ( or EOF, got %s", val.info()[1:])
	}
	n.c[0] = &Node{val, nil}
	val, err = match('I')
	if err != nil {
		return err
	}
	n.c[1] = &Node{val, nil}
	val, err = match('S')
	if err != nil {
		return err
	}
	if val.info()[1] != '{' {
		return fmt.Errorf("Expected {, got %s", val.info()[1:])
	}
	n.c[2] = &Node{val, nil}
	err = block(n.c[3])
	if err != nil {
		return err
	}
	val, err = match('S')
	if err != nil {
		return err
	}
	if val.info()[1] != '}' {
		return fmt.Errorf("Expected }, got %s", val.info()[1:])
	}
	n.c[4] = &Node{val, nil}
	val, err = match('S')
	if err != nil {
		return err
	}
	if val.info()[1] != ')' {
		return fmt.Errorf("Expected ), got %s", val.info()[1:])
	}
	n.c[5] = &Node{val, nil}
	return program(n.c[6])
}

func block(n *Node) error {
	n.v = BLOCK
	if lex[lex_i] == nil || lex[lex_i].info()[1] == '}' {
		return nil
	}
	n.c = []*Node{{}, {}}
	err := element(n.c[0])
	if err != nil {
		val, err := match('S')
		if err != nil {
			return err
		}
		if val.info()[1] != '{' {
			fmt.Println([]int{}[0])
			return fmt.Errorf("Expected ELEMENT or {, got %s", val.info()[1:])
		}
		n.c = []*Node{{val, nil}, {}, {}, {}}
		err = block(n.c[1])
		if err != nil {
			return err
		}
		val, err = match('S')
		if err != nil {
			return err
		}
		if val.info()[1] != '}' {
			return fmt.Errorf("Expected }, got %s", val.info()[1:])
		}
		n.c[2] = &Node{val, nil}
		return block(n.c[3])
	}
	return block(n.c[1])
}

func element(n *Node) error {
	n.v = ELEMENT
	val, err := match('S')
	if err != nil {
		t := lex[lex_i]
		if t == nil {
			return err
		}
		move()
		if t.info()[0] == 'I' {
			val, err = match('S')
			if err == nil {
				if val.info()[1] == '(' {
					n.c = []*Node{{t, nil}, {val, nil}, {lex[lex_i], nil}}
					move()
					return nil
				} else {
					lex_i--
				}
			}
		}
		n.c = []*Node{{t, nil}}
		return nil
	}
	if val.info()[1] != '[' {
		lex_i--
		return fmt.Errorf("Expected [, Identity or Number, got %s", val.info()[1:])
	}
	n.c = []*Node{{val, nil}, {}, {}}
	comment(n.c[1])
	val, err = match('S')
	if err != nil {
		return err
	}
	if val.info()[1] != ']' {
		return fmt.Errorf("Expected ], got %s", val.info()[1:])
	}
	*n.c[2] = Node{val, nil}
	return nil
}

func comment(n *Node) {
	n.v = COMMENT
	if lex[lex_i].info()[1] == ']' {
		return
	}
	n.c = []*Node{{lex[lex_i], nil}, {}}
	move()
	comment(n.c[1])
}

func parser(_tks []Infoer) (*Node, error) {
	lex = _tks
	lex = append(lex, nil)
	n := Node{}
	err := program(&n)
	if err != nil {
		return nil, err
	}
	return &n, nil
}
