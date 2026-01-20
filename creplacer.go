package main

func commentFilter(n *Node) *Node {
	if n.c == nil {
		return n
	}
	if n.c[0].c[0].v.get() == "S[" {
		return commentFilter(n.c[1])
	}
	n.c[1] = commentFilter(n.c[1])
	return n
}

func replacer(n *Node) *Node {
	return commentFilter(n)
}
