package main

import (
	"fmt"
	"os"
)

func output(n *Node, k uint32) {
	for range k {
		fmt.Print(" ")
	}
	fmt.Println(n.v.get())
	for range k {
		fmt.Print(" ")
	}
	fmt.Println("(")
	for _, i := range n.c {
		output(i, k+4)
	}
	for range k {
		fmt.Print(" ")
	}
	fmt.Println(")")
}

func main() {
	arg := os.Args
	if len(arg) < 2 {
		fmt.Println("Too few args")
		return
	}
	s := arg[1]
	info := false
	for i := 2; i < len(arg); i++ {
		if arg[i] == "--info" {
			info = true
		} else {
			fmt.Printf("Unknown flag: %s\n", arg[i])
			return
		}
	}
	if info {
		fmt.Println("Lexer:")
	}
	tks, err := lexer(s + ".cl")
	if err != nil {
		fmt.Println(err)
		return
	}
	if info {
		for _, i := range tks {
			fmt.Println(i.info())
		}
		fmt.Println("Parser:")
	}
	pst, err := parser(tks)
	if err != nil {
		fmt.Println(err)
		return
	}
	if info {
		output(pst, uint32(0))
		fmt.Println("Replacer:")
	}
	pst = replacer(pst)
	if info {
		output(pst, uint32(0))
		fmt.Println("Analyzer:")
	}
	err = analyzer(pst, s)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if info {
		fmt.Printf("%s.cl has been compiled into %s.s", s, s)
	}
}
