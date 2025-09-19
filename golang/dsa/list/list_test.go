package list

import (
	"testing"
)

type tStruct struct {
	Number int
}

func checkLen[T any](t *testing.T, l *LinkedList[T], len int) bool {
	if l.len != len {
		t.Errorf("\ncheckLen:\nl.len = %d\nwant %d", l.len, len)
		return false
	}
	return true
}

func checkPtrs[T any](t *testing.T, l *LinkedList[T], nodes []*Node[T]) bool {
	if !checkLen(t, l, len(nodes)) {
		return false
	}

	if len(nodes) == 0 {
		if l.head != nil || l.tail != nil {
			t.Errorf("\ncheckPtrs:\nl.head = %p, l.tail = %p\nboth should be nil", l.head, l.tail)
			return false
		}
		return true
	}

	if l.head != nodes[0] {
		t.Errorf("\ncheckPtrs:\nl.head = %+v\nshould be %+v", l.head, nodes[0])
		return false
	}
	if l.tail != nodes[len(nodes)-1] {
		t.Errorf("\ncheckPtrs:\nl.tail = %+v\nshould be %+v", l.head, nodes[len(nodes)-1])
		return false
	}

	if len(nodes) == 1 {
		return true
	}

	i := 0
	cur := l.head
	for cur.next != nil {
		i += 1
		if i == len(nodes)-1 {
			break
		}

		cur = cur.next
		if cur != nodes[i] {
			t.Errorf("\ncheckPtrs:\nnode[%d] = %+v\nshould be %+v", i, cur, nodes[i])
			return false
		}
	}

	return true
}

func checkSingly[T any](t *testing.T, l *LinkedList[T], nodes []*Node[T]) bool {
	if !checkLen(t, l, len(nodes)) {
		return false
	}

	if len(nodes) == 0 {
		return true
	}

	if l.head == nil && l.tail == nil {
		return true
	}

	if l.double {
		t.Error("l.double = true; should be false")
		return false
	}

	nextCount := 0
	cur := l.head
	for cur.next != nil {
		if cur.prev != nil {
			t.Errorf("cur.prev = %p; should be nil", cur.prev)
			return false
		}
		nextCount += 1
		cur = cur.next
		if l.circular && cur == l.head {
			break
		}
	}

	if l.circular {
		if nextCount != len(nodes) {
			t.Errorf("nextCount = %d; should be %d", nextCount, len(nodes))
			return false
		}
	} else {
		if nextCount != len(nodes)-1 {
			t.Errorf("nextCount = %d; should be %d", nextCount, len(nodes)-1)
			return false
		}
	}

	return true
}

func checkDoubly[T any](t *testing.T, l *LinkedList[T], nodes []*Node[T]) bool {
	if !checkLen(t, l, len(nodes)) {
		return false
	}

	if len(nodes) == 0 {
		return true
	}

	if l.head == nil && l.tail == nil {
		return true
	}

	if !l.double {
		t.Error("l.double = false; should be true")
		return false
	}

	nextCount := 0
	prevCount := 0
	cur := l.head
	for cur.next != nil {
		nextCount += 1
		if cur.prev != nil {
			prevCount += 1
		}
		cur = cur.next
		if !l.circular && cur.next == nil {
			if cur.prev != nil {
				prevCount += 1
			}
		}
		if l.circular && cur == l.head {
			break
		}
	}

	if l.circular {
		if nextCount != len(nodes) {
			t.Errorf("nextCount = %d; should be %d", nextCount, len(nodes))
			return false
		}
		if prevCount != len(nodes) {
			t.Errorf("prevCount = %d; should be %d", prevCount, len(nodes))
			return false
		}
	} else {
		if nextCount != len(nodes)-1 {
			t.Errorf("nextCount = %d; should be %d", nextCount, len(nodes)-1)
			return false
		}
		if prevCount != len(nodes)-1 {
			t.Errorf("prevCount = %d; should be %d", prevCount, len(nodes)-1)
			return false
		}
	}

	return true
}

func checkCircular[T any](t *testing.T, l *LinkedList[T]) bool {
	if !l.circular {
		t.Error("l.circular = false; should be true")
		return false
	}

	if l.tail.next != l.head {
		t.Errorf("l.tail.next = %p; should be %p", l.tail.next, l.head)
		return false
	}

	if l.double && l.head.prev != l.tail {
		t.Errorf("l.head.prev = %p; should be %p", l.head.prev, l.tail)
		return false
	}

	return true
}

func TestList(t *testing.T) {
	// check singly linked list

	l := New[tStruct]()
	checkPtrs(t, l, []*Node[tStruct]{})

	n1 := l.PrependValue(tStruct{Number: 1})
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.MoveToFront(n1)
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.MoveToBack(n1)
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.Remove(n1)
	checkPtrs(t, l, []*Node[tStruct]{})

	n2 := l.PrependValue(tStruct{Number: 2})
	n1 = l.PrependValue(tStruct{Number: 1})

	n3 := l.AppendValue(tStruct{Number: 3})
	n4 := l.AppendValue(tStruct{Number: 4})

	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n3, n4})
	checkSingly(t, l, []*Node[tStruct]{n1, n2, n3, n4})

	l.Remove(n2)
	checkPtrs(t, l, []*Node[tStruct]{n1, n3, n4})
	checkSingly(t, l, []*Node[tStruct]{n1, n3, n4})

	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkSingly(t, l, []*Node[tStruct]{n3, n1, n4})

	l.MoveToFront(n1)
	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkSingly(t, l, []*Node[tStruct]{n1, n4, n3})

	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkSingly(t, l, []*Node[tStruct]{n3, n1, n4})
	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkSingly(t, l, []*Node[tStruct]{n3, n1, n4})

	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkSingly(t, l, []*Node[tStruct]{n1, n4, n3})
	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkSingly(t, l, []*Node[tStruct]{n1, n4, n3})

	// insert before front
	n2 = l.InsertBefore(tStruct{Number: 2}, n1)
	checkPtrs(t, l, []*Node[tStruct]{n2, n1, n4, n3})
	l.Remove(n2)
	// insert before middle
	n2 = l.InsertBefore(tStruct{Number: 2}, n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert before back
	n2 = l.InsertBefore(tStruct{Number: 2}, n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)

	// insert after front
	n2 = l.InsertAfter(tStruct{Number: 2}, n1)
	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert after middle
	n2 = l.InsertAfter(tStruct{Number: 2}, n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)

	// insert after back
	n2 = l.InsertAfter(tStruct{Number: 2}, n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3, n2})

	if n := FindNodeByValue(l, n2.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n2)
	}

	if n := FindNodeByValue(l, n3.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n3)
	}

	if n := FindNodeByValue(l, n4.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n4)
	}

	if n := FindNodeByValue(l, n1.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n1)
	}

	l = l.MakeCircular()
	checkCircular(t, l)
	l.Remove(n2)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkSingly(t, l, []*Node[tStruct]{n1, n4, n3})
	checkCircular(t, l)

	l.Remove(n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n3})
	checkSingly(t, l, []*Node[tStruct]{n1, n3})
	checkCircular(t, l)

	l.Remove(n1)
	checkPtrs(t, l, []*Node[tStruct]{n3})
	checkSingly(t, l, []*Node[tStruct]{n3})
	checkCircular(t, l)

	// check doubly linked list

	l = l.Reset()
	checkPtrs(t, l, []*Node[tStruct]{})

	l = l.MakeDoubly()
	checkDoubly(t, l, []*Node[tStruct]{})

	n1 = l.PrependValue(tStruct{Number: 1})
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.MoveToFront(n1)
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.MoveToBack(n1)
	checkPtrs(t, l, []*Node[tStruct]{n1})

	l.Remove(n1)
	checkPtrs(t, l, []*Node[tStruct]{})

	n2 = l.PrependValue(tStruct{Number: 2})
	n1 = l.PrependValue(tStruct{Number: 1})

	n3 = l.AppendValue(tStruct{Number: 3})
	n4 = l.AppendValue(tStruct{Number: 4})

	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n3, n4})
	checkDoubly(t, l, []*Node[tStruct]{n1, n2, n3, n4})

	l.Remove(n2)
	checkPtrs(t, l, []*Node[tStruct]{n1, n3, n4})
	checkDoubly(t, l, []*Node[tStruct]{n1, n3, n4})

	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkDoubly(t, l, []*Node[tStruct]{n3, n1, n4})

	l.MoveToFront(n1)
	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkDoubly(t, l, []*Node[tStruct]{n1, n4, n3})

	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkDoubly(t, l, []*Node[tStruct]{n3, n1, n4})
	l.MoveToFront(n3)
	checkPtrs(t, l, []*Node[tStruct]{n3, n1, n4})
	checkDoubly(t, l, []*Node[tStruct]{n3, n1, n4})

	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkDoubly(t, l, []*Node[tStruct]{n1, n4, n3})
	l.MoveToBack(n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkDoubly(t, l, []*Node[tStruct]{n1, n4, n3})

	// insert before front
	n2 = l.InsertBefore(tStruct{Number: 2}, n1)
	checkPtrs(t, l, []*Node[tStruct]{n2, n1, n4, n3})
	l.Remove(n2)
	// insert before middle
	n2 = l.InsertBefore(tStruct{Number: 2}, n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert before back
	n2 = l.InsertBefore(tStruct{Number: 2}, n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)

	// insert after front
	n2 = l.InsertAfter(tStruct{Number: 2}, n1)
	checkPtrs(t, l, []*Node[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert after middle
	n2 = l.InsertAfter(tStruct{Number: 2}, n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)

	// insert after back
	n2 = l.InsertAfter(tStruct{Number: 2}, n3)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3, n2})

	if n := FindNodeByValue(l, n2.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n2)
	}

	if n := FindNodeByValue(l, n3.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n3)
	}

	if n := FindNodeByValue(l, n4.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n4)
	}

	if n := FindNodeByValue(l, n1.Value.Number, func(node *Node[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n1)
	}

	l = l.MakeCircular()
	checkCircular(t, l)
	l.Remove(n2)
	checkPtrs(t, l, []*Node[tStruct]{n1, n4, n3})
	checkDoubly(t, l, []*Node[tStruct]{n1, n4, n3})
	checkCircular(t, l)

	l.Remove(n4)
	checkPtrs(t, l, []*Node[tStruct]{n1, n3})
	checkDoubly(t, l, []*Node[tStruct]{n1, n3})
	checkCircular(t, l)

	l.Remove(n1)
	checkPtrs(t, l, []*Node[tStruct]{n3})
	checkDoubly(t, l, []*Node[tStruct]{n3})
	checkCircular(t, l)
}
