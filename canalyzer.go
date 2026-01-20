package main

import (
	"fmt"
	"os"
)

func printCall(asm *os.File, n *Node, k uint32) {
	if n.c == nil {
		return
	}
	fmt.Fprintln(asm, "    pushq %rbp")
	fmt.Fprintln(asm, "    movq %rsp, %rbp")
	fmt.Fprintln(asm, "    subq $32, %rsp")
	fmt.Fprintf(asm, "    leaq msg%d(%%rip), %%rcx\n", k)
	fmt.Fprintln(asm, "    call puts")
	fmt.Fprintln(asm, "    addq $32, %rsp")
	fmt.Fprintln(asm, "    popq %rbp")
	printCall(asm, n.c[1], k+1)
}

func printMsg(asm *os.File, n *Node, k uint32) {
	if n.c == nil {
		return
	}
	fmt.Fprintf(asm, "msg%d:\n", k)
	fmt.Fprintf(asm, "    .asciz \"%s\"\n", n.c[0].c[0].v.get()[1:])
	printMsg(asm, n.c[1], k+1)
}

func analyzer(n *Node, s string) error {
	asm, err := os.Create(s + ".s")
	if err != nil {
		return err
	}
	defer asm.Close()
	fmt.Fprintln(asm, ".global main")
	fmt.Fprintln(asm, ".extern puts")
	fmt.Fprintln(asm, "main:")
	printCall(asm, n, uint32(0))
	fmt.Fprintln(asm, "    ret")
	printMsg(asm, n, uint32(0))
	asm.Close()
	return nil
}
