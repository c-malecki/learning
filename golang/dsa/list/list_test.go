package list

import (
	"testing"
)

type tStruct struct {
	Value int
}

var testInts = []int{1, 2, 3, 4, 5, 6}
var testStructs = []tStruct{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}, {Value: 6}, {Value: 7}, {Value: 8}, {Value: 9}}

// func TestList(t *testing.T) {
// 	l := New[tStruct]()
// 	fmt.Print("New List\n")
// 	l.Print()

// 	fmt.Print("MakeCircular\n")
// 	l.MakeCircular()
// 	l.Print()

// 	fmt.Print("Append Nodes\n")
// 	for _, v := range testStructs {
// 		l.AppendValue(v)
// 	}
// 	l.Print()
// 	l.PrintNodeForward()
// 	l.PrintNodeReverse()

// 	fmt.Print("MakeDoubly\n")
// 	l.MakeDoubly()
// 	l.Print()
// 	l.PrintNodeForward()
// 	l.PrintNodeReverse()

// 	fmt.Print("MakeSingly\n")
// 	l.MakeSingly()
// 	l.Print()
// 	l.PrintNodeForward()
// 	l.PrintNodeReverse()

// 	fmt.Print("MakeLinear\n")
// 	l.MakeLinear()
// 	l.Print()
// 	l.PrintNodeForward()
// 	l.PrintNodeReverse()
// }

func TestMoves(t *testing.T) {
	l := New[int]()
	for i, v := range testInts {
		if i == 3 {
			break
		}
		l.AppendValue(v)
	}
	// l.MakeCircular()
	l.MakeDoubly()
	l.PrintNodeForward()

	node := FindNodeByValue(l, 2, func(node *Node[int]) int {
		return node.Value
	})

	l.MoveToFront(node)
	l.PrintNodeForward()

	l.MoveToBack(node)
	l.PrintNodeForward()
}

// func TestRemove(t *testing.T) {
// 	l := New[int]()
// 	for i, v := range testInts {
// 		if i == 3 {
// 			break
// 		}
// 		l.AppendValue(v)
// 	}
// 	l.MakeDoubly()
// 	l.MakeCircular()
// 	l.PrintNodeForward()
// 	l.Print()

// 	node := FindNodeByValue(l, 1, func(node *Node[int]) int {
// 		return node.Value
// 	})
// 	fmt.Printf("Remove %v\n", node.Value)
// 	l.Remove(node)
// 	l.PrintNodeForward()
// 	l.Print()

// 	node = FindNodeByValue(l, 3, func(node *Node[int]) int {
// 		return node.Value
// 	})
// 	fmt.Printf("Remove %v\n", node.Value)
// 	l.Remove(node)
// 	l.PrintNodeForward()
// 	l.Print()

// 	node = FindNodeByValue(l, 2, func(node *Node[int]) int {
// 		return node.Value
// 	})
// 	fmt.Printf("Remove %v\n", node.Value)
// 	l.Remove(node)
// 	l.PrintNodeForward()
// 	l.Print()
// }
