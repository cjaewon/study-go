# ⛓️ 링크드 리스트
![linked-list](https://user-images.githubusercontent.com/32125218/94662161-fb134100-0342-11eb-8e0a-28db316ec2b1.png)


**root**: 시작 노드를 가르키는 포인터
**Node** 링크드 리스트에서 연결되는 요소들

```go
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
```

더 성능 향상을 원하면 더블 링크드 리스트를 사용하면 된다.