package main

import (
	"fmt"
	"os"
)

var kc, km uint32
var asm *os.File

func printCall(n *Node) {
	if n.c == nil {
		return
	}
	if len(n.c) != 2 {
		if len(n.c) == 4 {
			printCall(n.c[1])
			printCall(n.c[3])
		}
		return
	}
	if len(n.c[0].c) != 3 {
		fmt.Fprintln(asm, "    pushq %rbp")
		fmt.Fprintln(asm, "    movq %rsp, %rbp")
		fmt.Fprintln(asm, "    subq $32, %rsp")
		fmt.Fprintf(asm, "    leaq msg%d(%%rip), %%rcx\n", kc)
		fmt.Fprintln(asm, "    call puts")
		fmt.Fprintln(asm, "    addq $32, %rsp")
		fmt.Fprintln(asm, "    popq %rbp")
		kc++
	} else {
		fmt.Fprintf(asm, "    call %s\n", n.c[0].c[0].v.get()[1:])
	}
	printCall(n.c[1])
	if len(n.c) == 4 {
		printCall(n.c[3])
	}
}

func printMsg(n *Node) {
	if n.c == nil {
		return
	}
	if len(n.c) != 2 {
		if len(n.c) == 4 {
			printMsg(n.c[1])
			printMsg(n.c[3])
		}
		return
	}
	if len(n.c[0].c) != 3 {
		fmt.Fprintf(asm, "msg%d:\n", km)
		fmt.Fprintf(asm, "    .asciz \"%s\"\n", n.c[0].c[0].v.get()[1:])
		km++
	}
	printMsg(n.c[1])
}

func printFunc(n *Node) {
	if n.c == nil {
		return
	}
	fmt.Fprintf(asm, "%s:\n", n.c[1].v.get()[1:])
	fmt.Fprintln(asm, "    pushq %rbp")
	fmt.Fprintln(asm, "    movq %rsp, %rbp")
	printCall(n.c[3])
	fmt.Fprintln(asm, "    popq %rbp")
	fmt.Fprintln(asm, "    ret")
	printMsg(n.c[3])
	printFunc(n.c[6])
}

func analyzer(n *Node, s string) error {
	var err error
	asm, err = os.Create(s + ".s")
	if err != nil {
		return err
	}
	defer asm.Close()
	fmt.Fprintln(asm, ".global main")
	fmt.Fprintln(asm, ".extern puts")
	printFunc(n)
	return nil
}
