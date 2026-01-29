package main

func commentFilter(n *Node) *Node {
	if n.c == nil {
		return n
	}
	if len(n.c) == 2 && n.c[0].c[0].v.get()[1] == '[' {
		return commentFilter(n.c[1])
	}
	n.c[1] = commentFilter(n.c[1])
	if len(n.c) == 4 {
		n.c[3] = commentFilter(n.c[3])
	}
	return n
}

func replacer(n *Node) *Node {
	if len(n.c) == 7 {
		n.c[3] = commentFilter(n.c[3])
		n.c[6] = replacer(n.c[6])
	}
	return n
}
