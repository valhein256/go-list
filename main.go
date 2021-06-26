package main

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Add(value int) {
	node := &Node{value: value}
	if l.First() == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

type Dlist struct {
	head *Node
	tail *Node
}

func (dl *Dlist) First() *Node {
	return dl.head
}
func (dl *Dlist) Last() *Node {
	return dl.tail
}

func (dl *Dlist) Add(value int) {
	node := &Node{value: value}
	if dl.First() == nil {
		dl.head = node
	} else {
		dl.tail.next = node
		node.prev = dl.tail
	}
	dl.tail = node
}

func main() {
	l := &List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	n := l.First()
	println("First item in List: %d", l.First().value)
	println("First item in List: %d", l.Last().value)
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}

	dl := &Dlist{}
	dl.Add(-1)
	dl.Add(-2)
	dl.Add(-3)
	dl.Add(-4)
	m := dl.Last()
	println("First item in List: %d", dl.First().value)
	println("First item in List: %d", dl.Last().value)
	for {
		println(m.value)
		m = m.Prev()
		if m == nil {
			break
		}
	}
}
