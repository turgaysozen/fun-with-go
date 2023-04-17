package main

import "fmt"

type Node struct {
	left, right *Node
	data        int
}

func (n *Node) insert(data int) {
	if n == nil {
		return
	}

	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.insert(data)
		}
	}
}

func (n *Node) printNodes() {
	if n == nil {
		return
	}

	if n.left != nil {
		n.left.printNodes()
	}
	fmt.Println("Node: ", n.data)
	if n.right != nil {
		n.right.printNodes()
	}
}

func (n *Node) inOrderTraversal() {
	if n == nil {
		return
	}
	n.left.inOrderTraversal()
	fmt.Printf("%d ", n.data)
	n.right.inOrderTraversal()
}

func (n *Node) search(data int) *Node {
	if n == nil {
		return nil
	}
	if n.data == data {
		return n
	}
	if data < n.data {
		return n.left.search(data)
	} else {
		return n.right.search(data)
	}
}
