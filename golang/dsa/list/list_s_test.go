package list

import (
	"testing"
)

func checkSinglyLen[T any](t *testing.T, l *SinglyLinkedList[T], len int) bool {
	if l.len != len {
		t.Errorf("\n=== checkLen ===\nl.len = %d\nwant %d", l.len, len)
		return false
	}
	return true
}

func checkSinglyPtrs[T any](t *testing.T, l *SinglyLinkedList[T], nodes []*NodeS[T]) bool {
	if !checkSinglyLen(t, l, len(nodes)) {
		return false
	}

	if len(nodes) == 0 {
		if l.head != nil || l.tail != nil {
			t.Errorf("\n=== checkSinglyPtrs ===\nl.head = %p %+v\nl.tail = %p %+v\nboth should be nil", l.head, l.head, l.tail, l.tail)
			return false
		}
		return true
	}

	if l.head != nodes[0] {
		t.Errorf("\n=== checkSinglyPtrs ===\nl.head = %p %+v\nshould be %p %+v", l.head, l.head, nodes[0], nodes[0])
		return false
	}
	if l.tail != nodes[len(nodes)-1] {
		t.Errorf("\n=== checkSinglyPtrs ===\nl.tail = %p %+v\nshould be %p %+v", l.head, l.head, nodes[len(nodes)-1], nodes[len(nodes)-1])
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
			t.Errorf("\n=== checkSinglyPtrs ===\nnode[%d] = %p %+v\nshould be %p %+v", i, cur, cur, nodes[i], nodes[i])
			return false
		}
	}

	return true
}

func checkSinglyNextCount[T any](t *testing.T, l *SinglyLinkedList[T], nodes []*NodeS[T]) bool {
	if !checkSinglyLen(t, l, len(nodes)) {
		return false
	}

	if len(nodes) == 0 {
		return true
	}

	if l.head == nil && l.tail == nil {
		return true
	}

	nextCount := 0
	cur := l.head
	for cur.next != nil {
		nextCount += 1
		cur = cur.next
		if l.circular && cur == l.head {
			break
		}
	}

	if l.circular {
		if nextCount != len(nodes) {
			t.Errorf("\n=== checkSinglyNextCount ===\nnextCount = %d; should be %d", nextCount, len(nodes))
			return false
		}
	} else {
		if nextCount != len(nodes)-1 {
			t.Errorf("\n=== checkSinglyNextCount ===\nnextCount = %d; should be %d", nextCount, len(nodes)-1)
			return false
		}
	}

	return true
}

func checkSinglyCircular[T any](t *testing.T, l *SinglyLinkedList[T]) bool {
	if !l.circular {
		t.Error("\n=== checkCircular ===\nl.circular = false; should be true")
		return false
	}

	if l.tail.next != l.head {
		t.Errorf("\n=== checkCircular ===\nl.tail.next = %p %+v; should be %p %+v", l.tail.next, l.tail.next, l.head, l.head)
		return false
	}

	return true
}

func TestSinglyLinkedList(t *testing.T) {
	l := NewSinglyLinkedList[tStruct]()
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{})

	n1 := l.PrependValue(tStruct{Number: 1})
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1})

	l.MoveToFront(n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1})

	l.MoveToBack(n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1})

	l.Remove(n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{})

	n2 := l.PrependValue(tStruct{Number: 2})
	n1 = l.PrependValue(tStruct{Number: 1})

	n3 := l.AppendValue(tStruct{Number: 3})
	n4 := l.AppendValue(tStruct{Number: 4})

	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n2, n3, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n2, n3, n4})

	l.Remove(n2)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n3, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n3, n4})

	l.MoveToFront(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n3, n1, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n3, n1, n4})

	l.MoveToFront(n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n3, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n3, n4})

	l.MoveToBack(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n4, n3})

	l.MoveToFront(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n3, n1, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n3, n1, n4})
	l.MoveToFront(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n3, n1, n4})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n3, n1, n4})

	l.MoveToBack(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n4, n3})
	l.MoveToBack(n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n4, n3})

	// insert before front
	n2 = l.InsertBefore(tStruct{Number: 2}, n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n2, n1, n4, n3})
	l.Remove(n2)
	// insert before middle
	n2 = l.InsertBefore(tStruct{Number: 2}, n4)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert before back
	n2 = l.InsertBefore(tStruct{Number: 2}, n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)

	// insert after front
	n2 = l.InsertAfter(tStruct{Number: 2}, n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n2, n4, n3})
	l.Remove(n2)
	// insert after middle
	n2 = l.InsertAfter(tStruct{Number: 2}, n4)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n2, n3})
	l.Remove(n2)
	// insert after back
	n2 = l.InsertAfter(tStruct{Number: 2}, n3)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n3, n2})

	l.MakeCircular()
	checkSinglyCircular(t, l)
	l.Remove(n2)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n4, n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n4, n3})
	checkSinglyCircular(t, l)

	l.Remove(n4)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n1, n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n1, n3})
	checkSinglyCircular(t, l)

	l.Remove(n1)
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{n3})
	checkSinglyNextCount(t, l, []*NodeS[tStruct]{n3})
	checkSinglyCircular(t, l)
}
