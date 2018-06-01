package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// DeleteDuplication deletes the duplicated nodes of the sorted linkedlist
func DeleteDuplication(head **datastructure.Node) {
	if head == nil || *head == nil {
		return
	}
	var pre *datastructure.Node
	node := *head
	for node != nil && node.Next != nil {
		if datastructure.Equal(node.Data, node.Next.Data) {
			value := node.Data
			for node != nil && datastructure.Equal(value, node.Data) {
				toBeDel := node
				node = node.Next
				// delete node
				toBeDel.Data = nil
				toBeDel.Next = nil
				toBeDel = nil
			}
			if pre == nil {
				*head = node
			} else {
				pre.Next = node
			}
		} else {
			pre = node
			node = node.Next
		}
	}
}
