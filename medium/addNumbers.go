package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	head := l3
	temp := 0

	for l1 != nil || l2 != nil || temp > 0 {
		l3.Next = new(ListNode)
		l3 = l3.Next
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		l3.Val = temp % 10
		temp /= 10
	}
	return head.Next
}

/**
 *链表逆打印
 */

func reversePrint(head *ListNode) []int {
	var list []int
	for head != nil {
		list = append([]int{head.Val}, list...)
		head = head.Next
	}
	return list
}
