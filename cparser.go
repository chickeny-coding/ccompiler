package main

import (
	"fmt"
)

type Symbol byte

const (
	_ Symbol = iota
	PROGRAM
	BLOCK
	COMMENT
)

var ss []string = []string{
	"",
	"PROGRAM",
	"BLOCK",
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

func current() Infoer {
	if lex[lex_i] == nil {
		return nil
	}
	return lex[lex_i]
}

func move() {
	lex_i++
}

func match(c byte) (Infoer, error) {
	if lex[lex_i] == nil {
		return nil, fmt.Errorf("Reached EOF")
	}
	if current().info()[0] != c {
		return nil, fmt.Errorf("Expected %s, got %s", itm[c], current().info())
	}
	v := lex[lex_i]
	move()
	return v, nil
}

func program(n *Node) error {
	n.v = PROGRAM
	n.c = []*Node{{}, {}}
	err := block(n.c[0])
	if err != nil {
		return err
	}
	if current() != nil {
		return program(n.c[1])
	}
	n.c[1].v = PROGRAM
	n.c[1].c = nil
	return nil
}

func block(n *Node) error {
	n.v = BLOCK
	val, err := match('S')
	if err != nil {
		n.c = []*Node{{current(), nil}}
		move()
		return nil
	}
	if (val.info()[1] != '[') {
		return fmt.Errorf("Expected [, Identity or Number, got %s", current().info())
	}
	n.c = []*Node{{val, nil}, {}, {}}
	comment(n.c[1])
	val, err = match('S')
	if err != nil || val.info()[1] != ']' {
		return fmt.Errorf("Expected ], got %s", current().info())
	}
	*n.c[2] = Node{val, nil}
	return nil
}

func comment(n *Node) {
	n.v = COMMENT
	if current().info()[1] == ']' {
		return
	}
	n.c = []*Node{{current(), nil}, {}}
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
