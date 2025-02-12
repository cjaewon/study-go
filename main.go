package main

import (
	"fmt"
)

// Node :
type Node struct {
	next *Node // 다음 노드를 가르킴
	val  int
}

// Add :
func Add(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

// Remove :
func Remove(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next
		return root, tail
	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}

	if node == tail {
		prev.next = nil
		tail = prev
	} else {
		prev.next = prev.next.next // 가비지 컬렉터가 자동으로 제거
	}

	return root, tail
}

// PrintNodes :
func PrintNodes(root *Node) {
	n := root
	for {
		fmt.Print(n.val, " ")
		if n.next == nil {
			break
		}

		n = n.next
	}

}

func main() {
	var root *Node = &Node{
		val: 0,
	}
	var tail *Node = root

	for i := 1; i <= 10; i++ {
		tail = Add(tail, i)
	}

	root, tail = Remove(root.next, root, tail)

	PrintNodes(root)
}
