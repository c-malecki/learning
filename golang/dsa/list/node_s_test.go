package list

import "testing"

func TestNodeS(t *testing.T) {
	l := NewSinglyLinkedList[tStruct]()

	n1 := l.AppendValue(tStruct{Number: 1})
	n2 := l.AppendValue(tStruct{Number: 2})
	n3 := l.AppendValue(tStruct{Number: 3})
	n4 := l.AppendValue(tStruct{Number: 4})

	if n := FindNodeSByValue(l, n1.Value.Number, func(node *NodeS[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n1)
	}

	if n := FindNodeSByValue(l, n2.Value.Number, func(node *NodeS[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n2)
	}

	if n := FindNodeSByValue(l, n3.Value.Number, func(node *NodeS[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n3)
	}

	if n := FindNodeSByValue(l, n4.Value.Number, func(node *NodeS[tStruct]) int {
		return node.Value.Number
	}); n == nil {
		t.Errorf("node = %+v; should be %+v", n, n4)
	}

	sum := 0
	for n := l.Front(); n != nil; n = n.Next() {
		sum += n.Value.Number
	}
	if sum != 10 {
		t.Errorf("sum = %d; should be 10", sum)
	}

	var next *NodeS[tStruct]
	for n := l.Front(); n != nil; n = next {
		next = n.Next()
		l.Remove(n)
	}
	checkSinglyPtrs(t, l, []*NodeS[tStruct]{})
}
