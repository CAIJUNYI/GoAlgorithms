package test

import (
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
	"github.com/CAIJUNYI/GoAlgorithms/interviews"
)

func Test04ReplaceBlank(t *testing.T) {
	s1 := " abc def  g  "
	r1 := interviews.ReplaceBlank(s1)
	exp1 := "%20abc%20def%20%20g%20%20"
	if r1 != exp1 {
		t.Errorf("expected %v and got %v", exp1, r1)
	}

	s2 := ""
	r2 := interviews.ReplaceBlank(s2)
	exp2 := ""
	if r2 != exp2 {
		t.Errorf("expected %v and got %v", exp2, r2)
	}

	s3 := "abc"
	r3 := interviews.ReplaceBlank(s3)
	exp3 := "abc"
	if r3 != exp3 {
		t.Errorf("expected %v and got %v", exp3, r3)
	}
}

func Test05PrintLinkedList(t *testing.T) {
	var ll datastructure.LinkedList
	ll.Append("1")
	ll.Append("2")
	ll.Append("3")
	ll.Append("4")

	interviews.PrintLinkedlistReversely(&ll)
}

func Test06ReconstructTree(t *testing.T) {
	preorder := []int{8, 4, 2, 1, 3, 6, 5, 7, 10, 9}
	inorder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := interviews.ReConstructBiTree(preorder, inorder)
	var result []string
	expPostOrder := []string{"1", "3", "2", "5", "7", "6", "4", "9", "10", "8"}

	visit := func(e datastructure.Elem) {
		result = append(result, fmt.Sprintf("%v", e))
	}

	tree.PostOrderTraverseIter(tree.Root, visit)
	if len(result) == 0 {
		t.Errorf("expected %v\n got %v", expPostOrder, result)
	}
	for i, e := range result {
		if expPostOrder[i] != e {
			t.Errorf("expected %v and got %v", expPostOrder[i], e)
		}
	}
}

func Test08MinOfRotatedArray(t *testing.T) {
	arr1 := []int{3, 4, 5, 1, 2}
	min1, ok := interviews.MinOfRotatedArray(arr1)
	if !ok || min1 != 1 {
		t.Errorf("expected 1 and got %v", min1)
	}

	arr2 := []int{1, 0, 1, 1, 1}
	min2, ok := interviews.MinOfRotatedArray(arr2)
	if !ok || min2 != 0 {
		t.Errorf("expected 0 and got %v", min2)
	}

	arr3 := []int{1, 1, 1, 0, 1}
	min3, ok := interviews.MinOfRotatedArray(arr3)
	if !ok || min3 != 0 {
		t.Errorf("expected 0 and got %v", min3)
	}

	arr4 := []int{1, 2, 3, 4, 5}
	min4, ok := interviews.MinOfRotatedArray(arr4)
	if !ok || min4 != 1 {
		t.Errorf("expected 1 and got %v", min4)
	}
	arr5 := []int{}
	min5, ok := interviews.MinOfRotatedArray(arr5)
	if ok || min5 != -1 {
		t.Errorf("expected -1 and got %v", min5)
	}
}

func Test09Fib(t *testing.T) {
	if n, ok := interviews.Fib(3); !ok || n != 2 {
		t.Errorf("expected 2 and got %v", n)
	}
}

func Test10NumOf1(t *testing.T) {
	cnt1 := interviews.NumOf1(1)
	if cnt1 != 1 {
		t.Errorf("expected 1 and got %v", cnt1)
	}

	cnt2 := interviews.NumOf1(0)
	if cnt2 != 0 {
		t.Errorf("expected 0 and got %v", cnt2)
	}

	cnt3 := interviews.NumOf1(-1)
	if cnt3 != 64 {
		t.Errorf("expected 64 and got %v", cnt3)
	}

}

func Test10IsPowerOf2(t *testing.T) {
	b1 := interviews.IsPowerOf2(4)
	if !b1 {
		t.Errorf("expected true and got %v", b1)
	}

	b2 := interviews.IsPowerOf2(0)
	if b2 {
		t.Errorf("expected false and got %v", b2)
	}

	b3 := interviews.IsPowerOf2(3)
	if b3 {
		t.Errorf("expected false and got %v", b3)
	}
}

func Test10NumOfChangingBits(t *testing.T) {
	if n1 := interviews.NumOfChangingBits(10, 13); n1 != 3 {
		t.Errorf("expected 3 and got %v", n1)
	}

	if n2 := interviews.NumOfChangingBits(-1, 1); n2 != 63 {
		t.Errorf("expected 63 and got %v", n2)
	}
}

func Test11Power(t *testing.T) {
	if n1, ok := interviews.Power(2.0, 2); !ok || !interviews.Equal(n1, 4) {
		t.Errorf("expected 4 and got %v", n1)
	}
	if n2, ok := interviews.Power(0, 2); !ok || !interviews.Equal(n2, 0) {
		t.Errorf("expected 0 and got %v", n2)
	}
	if n3, ok := interviews.Power(-2.0, 3); !ok || !interviews.Equal(n3, -8) {
		t.Errorf("expected -8 and got %v", n3)
	}

	if n4, ok := interviews.Power(-2.0, -2); !ok || !interviews.Equal(n4, 0.25) {
		t.Errorf("expected 0.25 and got %v", n4)
	}
}

func Test12Print1ToMax(t *testing.T) {
	interviews.Print1ToMaxOfNDigits(3)
}

func Test13DeleteNode(t *testing.T) {
	l1 := datastructure.Node{Data: 1, Next: nil}
	lP := &l1
	d1 := interviews.DeleteNode(&lP, lP)
	if *d1 != 1 || lP != nil {
		t.Errorf("expected 1 and got %v", *d1)
	}

	n3 := datastructure.Node{Data: 3, Next: nil}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}
	n1P := &n1
	if d2 := interviews.DeleteNode(&n1P, &n1); *d2 != 1 || n1P.Data != 2 {
		t.Errorf("expected 1 and got %v", *d2)
	}
	if d3 := interviews.DeleteNode(&n1P, &n3); *d3 != 3 || n1P.Data != 2 {
		t.Errorf("expected 3 and got %v", *d3)
	}

	if d := interviews.DeleteNode(nil, nil); d != nil {
		t.Errorf("expected nil and got %v", d)
	}
}

func Test14ReorderOddEven(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	seq := interviews.ReorderOddEven2(arr)
	exp := []int{1, 5, 3, 4, 2}
	for i, e := range seq {
		if e != exp[i] {
			t.Errorf("expected %v and got %v", exp[i], e)
		}
	}
}

func Test15FindKthToTail(t *testing.T) {
	n5 := datastructure.Node{Data: 5, Next: nil}
	n4 := datastructure.Node{Data: 4, Next: &n5}
	n3 := datastructure.Node{Data: 3, Next: &n4}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}

	if n := interviews.FindKthToTail(&n1, 1); (*n).Data != 5 {
		t.Errorf("expected 5 and got %v", (*n).Data)
	}
	if n := interviews.FindKthToTail(&n1, 5); (*n).Data != 1 {
		t.Errorf("expected 1 and got %v", (*n).Data)
	}
	if n := interviews.FindKthToTail(&n1, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

	if n := interviews.FindKthToTail(&n1, 6); n != nil {
		t.Errorf("expected nil and got %v", n)
	}
	if n := interviews.FindKthToTail(nil, 0); n != nil {
		t.Errorf("expected nil and got %v", n)
	}

}
func Test15FindMidNode(t *testing.T) {
	n5 := datastructure.Node{Data: 5, Next: nil}
	n4 := datastructure.Node{Data: 4, Next: &n5}
	n3 := datastructure.Node{Data: 3, Next: &n4}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}

	if n := interviews.FindMidNode(&n1); (*n).Data != 3 {
		t.Errorf("expected 3 and got %v", (*n).Data)
	}

	n6 := datastructure.Node{Data: 6, Next: nil}
	n5.Next = &n6
	if n := interviews.FindMidNode(&n1); (*n).Data != 3 {
		t.Errorf("expected 3 and got %v", (*n).Data)
	}
}

func Test15IsLoop(t *testing.T) {
	n5 := datastructure.Node{Data: 5, Next: nil}
	n4 := datastructure.Node{Data: 4, Next: &n5}
	n3 := datastructure.Node{Data: 3, Next: &n4}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}
	if b := interviews.IsLoop(&n1); b {
		t.Errorf("expected false and got %v", b)
	}

	n5.Next = &n1
	if b := interviews.IsLoop(&n1); !b {
		t.Errorf("expected true and got %v", b)
	}
}

func Test16ReverseLinkedlist(t *testing.T) {
	n3 := datastructure.Node{Data: 3, Next: nil}
	n2 := datastructure.Node{Data: 2, Next: &n3}
	n1 := datastructure.Node{Data: 1, Next: &n2}

	if n := interviews.ReverseLinkedlist(&n1); n != &n3 || n.Next != &n2 || n.Next.Next != &n1 {
		t.Errorf("failed reversing linkedlist")
	}

	if n := interviews.ReverseLinkedlist(&n1); n != &n1 || n.Next != nil {
		t.Errorf("failed reversing linkedlist")
	}

	if n := interviews.ReverseLinkedlist(nil); n != nil {
		t.Errorf("failed reversing linkedlist")
	}
}

func Test17MergeLinkedlist(t *testing.T) {
	n4 := datastructure.Node{Data: 5, Next: nil}
	n3 := datastructure.Node{Data: 4, Next: nil}
	n2 := datastructure.Node{Data: 3, Next: &n4}
	n1 := datastructure.Node{Data: 2, Next: &n3}

	if n := interviews.MergeLinkedlist(&n1, &n2); n != &n1 || n.Next != &n2 || n.Next.Next != &n3 || n.Next.Next.Next != &n4 {
		t.Errorf("failed merging linkedlist")
	}
}

func Test18HasSubtree(t *testing.T) {
	var t1 datastructure.BST
	t1.Insert(8, "8")
	t1.Insert(4, "4")
	t1.Insert(10, "10")
	t1.Insert(2, "2")
	t1.Insert(6, "6")
	t1.Insert(1, "1")

	var t2 datastructure.BST
	t2.Insert(10, "10")
	t2.Insert(2, "2")
	t2.Insert(6, "6")
	// t2 is a subtree of t1
	if has := interviews.HasSubtree(t1.Root, t2.Root); !has {
		t.Errorf("expected true and got %v", has)
	}
	// nil pointer
	if has := interviews.HasSubtree(nil, nil); has {
		t.Errorf("expected false and got %v", has)
	}

	var t3 datastructure.BST
	t3.Insert(1, "1")
	var t4 datastructure.BST
	t4.Insert(1, "1")
	// only one node
	if has := interviews.HasSubtree(t3.Root, t4.Root); !has {
		t.Errorf("expected true and got %v", has)
	}
	// t4 is not a subtree of t3
	t4.Insert(2, "2")
	if has := interviews.HasSubtree(t3.Root, t4.Root); has {
		t.Errorf("expected false and got %v", has)
	}
}

func Test19TreeMirror(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(8, "8")
	bst.Insert(4, "4")
	bst.Insert(10, "10")
	bst.Insert(2, "2")

	interviews.BiTreeMirror(bst.Root)
	if bst.Root.LChild.Key != 10 || bst.Root.RChild.Key != 4 || bst.Root.RChild.RChild.Key != 2 || bst.Root.RChild.LChild != nil {
		t.Errorf("failed mirror binary tree")
	}
}

func Test20VisitMatrixClockwisely(t *testing.T) {
	m1 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	result := []int{}
	visit := func(e int) {
		result = append(result, e)
	}

	interviews.VisitMatrixClockwisely(m1, 4, 4, visit)
	exp1 := []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}
	for i, e := range result {
		if e != exp1[i] {
			t.Errorf("expected %v and got %v", exp1[i], e)
		}
	}
	// only one row
	m2 := [][]int{
		{1, 2, 3, 4},
	}
	result = []int{}
	exp2 := []int{1, 2, 3, 4}
	interviews.VisitMatrixClockwisely(m2, 1, 4, visit)
	for i, e := range result {
		if e != exp2[i] {
			t.Errorf("expected %v and got %v", exp2[i], e)
		}
	}
	// only one col
	m3 := [][]int{
		{1},
		{4},
		{3},
		{2},
	}
	result = []int{}
	exp3 := []int{1, 4, 3, 2}
	interviews.VisitMatrixClockwisely(m3, 4, 1, visit)
	for i, e := range result {
		if e != exp3[i] {
			t.Errorf("expected %v and got %v", exp3[i], e)
		}
	}
	// only one entry
	m4 := [][]int{{3}}
	result = []int{}
	exp4 := []int{3}
	interviews.VisitMatrixClockwisely(m4, 1, 1, visit)
	for i, e := range result {
		if e != exp4[i] {
			t.Errorf("expected %v and got %v", exp4[i], e)
		}
	}
}

func Test21StackWithMin(t *testing.T) {
	s := interviews.StackWithMin{}
	s.Push(3)
	s.Push(4)
	s.Push(2)

	if m := s.Min(); (*m).(int) != 2 {
		t.Errorf("expected 2 and got %v", m)
	}
	s.Pop()
	s.Push(1)
	if m := s.Min(); (*m).(int) != 1 {
		t.Errorf("expected 1 and got %v", m)
	}
}

func Test22IsPopOrder(t *testing.T) {
	pushOrder := []int{1, 2, 3, 4, 5}
	popOrder := []int{3, 5, 4, 2, 1}

	if order := interviews.IsPopOrder(pushOrder, popOrder); !order {
		t.Errorf("expected true and got %v", order)
	}
}

func Test24VerifyPostSeqOfBST(t *testing.T) {
	arr1 := []int{5, 7, 6, 9, 11, 10, 8}
	if b := interviews.VerifyPostSeqOfBST(arr1, 7); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr2 := []int{4, 5, 8, 8}
	if b := interviews.VerifyPostSeqOfBST(arr2, 4); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr3 := []int{8, 9, 10}
	if b := interviews.VerifyPostSeqOfBST(arr3, 3); !b {
		t.Errorf("expected true and got %v", b)
	}
	arr4 := []int{8}
	if b := interviews.VerifyPostSeqOfBST(arr4, 1); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr5 := []int{}
	if b := interviews.VerifyPostSeqOfBST(arr5, 0); b {
		t.Errorf("expected false and got %v", b)
	}

	if b := interviews.VerifyPostSeqOfBST(nil, 0); b {
		t.Errorf("expected false and got %v", b)
	}
}

func Test24VerifyPreSeqOfBST(t *testing.T) {
	arr1 := []int{8, 6, 5, 7, 10, 9, 11}
	if b := interviews.VerifyPreSeqOfBST(arr1, 7); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr2 := []int{8, 8, 5, 4}
	if b := interviews.VerifyPreSeqOfBST(arr2, 4); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr3 := []int{8, 9, 10}
	if b := interviews.VerifyPreSeqOfBST(arr3, 3); !b {
		t.Errorf("expected true and got %v", b)
	}
	arr4 := []int{8}
	if b := interviews.VerifyPreSeqOfBST(arr4, 1); !b {
		t.Errorf("expected true and got %v", b)
	}

	arr5 := []int{}
	if b := interviews.VerifyPreSeqOfBST(arr5, 0); b {
		t.Errorf("expected false and got %v", b)
	}

	if b := interviews.VerifyPreSeqOfBST(nil, 0); b {
		t.Errorf("expected false and got %v", b)
	}
}

func Test25FindPathOfTree(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(10, "10")
	bst.Insert(5, "5")
	bst.Insert(4, "4")
	bst.Insert(7, "7")
	bst.Insert(12, "12")
	interviews.FindPathofTree(bst.Root, 22)

	t.Errorf("testing...")
}

func Test26CloneComplexLinkedlist(t *testing.T) {
	var a, b, c, d, e datastructure.ComplexListNode
	a.Data, a.Next, a.Sibling = "A", &b, &c
	b.Data, b.Next, b.Sibling = "B", &c, &e
	c.Data, c.Next, c.Sibling = "C", &d, nil
	d.Data, d.Next, d.Sibling = "D", &e, &b
	e.Data, e.Next, e.Sibling = "E", nil, nil

	clone := interviews.CloneComplexLinkedlist(&a)

	for node := &a; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}

	a.Data, a.Next, a.Sibling = "A", &b, &c
	b.Data, b.Next, b.Sibling = "B", &c, &e
	c.Data, c.Next, c.Sibling = "C", &d, &c
	d.Data, d.Next, d.Sibling = "D", &e, &b
	e.Data, e.Next, e.Sibling = "E", nil, &b

	clone = interviews.CloneComplexLinkedlist(&a)

	for node := &a; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}

	o := datastructure.ComplexListNode{Data: "O", Next: nil, Sibling: nil}
	clone = interviews.CloneComplexLinkedlist(&o)
	for node := &o; clone != nil && node != nil; clone, node = clone.Next, node.Next {
		if clone.Data != node.Data {
			t.Errorf("expected %v and got %v", node.Data, clone.Data)
		}
		if node.Sibling != nil {
			if clone.Sibling.Data != node.Sibling.Data {
				t.Errorf("expected %v and got %v", node.Data, clone.Data)
			}
		}
	}
}

func Test27ConvertBSTToDoubleLinkedlist(t *testing.T) {
	var bst datastructure.BST
	bst.Insert(10, "10")
	bst.Insert(6, "6")
	bst.Insert(14, "14")
	bst.Insert(4, "4")
	bst.Insert(8, "8")
	bst.Insert(12, "12")
	bst.Insert(16, "16")

	head := interviews.BSTToDoubleLinkedlist(bst.Root)
	if head.Key != 4 || head.LChild != nil || head.RChild.Key != 6 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	var bst2 datastructure.BST
	bst2.Insert(10, "10")
	head = interviews.BSTToDoubleLinkedlist(bst2.Root)
	if head.Key != 10 || head.LChild != nil || head.RChild != nil {
		t.Errorf("failed converting bst to double linkedlist")
	}

	bst2.Insert(6, "6")
	bst2.Insert(4, "4")
	head = interviews.BSTToDoubleLinkedlist(bst2.Root)
	if head.Key != 4 || head.LChild != nil || head.RChild.Key != 6 || head.RChild.RChild.Key != 10 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	var bst3 datastructure.BST
	bst3.Insert(10, "10")
	bst3.Insert(14, "14")
	bst3.Insert(16, "16")
	head = interviews.BSTToDoubleLinkedlist(bst3.Root)
	if head.Key != 10 || head.LChild != nil || head.RChild.Key != 14 || head.RChild.RChild.Key != 16 {
		t.Errorf("failed converting bst to double linkedlist")
	}

	head = interviews.BSTToDoubleLinkedlist(nil)
	if head != nil {
		t.Errorf("failed converting bst to double linkedlist")
	}
}

func Test40FindTwoNumsAppearOnce(t *testing.T) {
	arr1 := []int{2, 4, 3, 6, 3, 2, 5, 5}
	var n1, n2 int
	ok := interviews.FindTwoNumsAppearOnce(arr1, &n1, &n2)
	if !ok || n1 != 6 || n2 != 4 {
		t.Errorf("failed to find two numbers that appear once")
	}
}
