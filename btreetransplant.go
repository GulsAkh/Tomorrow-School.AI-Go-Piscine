package main

import "fmt"

type TreeNode struct {
	Data                string
	Left, Right, Parent *TreeNode
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Data == elem {
		return root
	}
	if elem > root.Data {
		if root.Right != nil {
			root.Right.Parent = root
		}
		return BTreeSearchItem(root.Right, elem)
	} else {
		if root.Left != nil {
			root.Left.Parent = root
		}
		return BTreeSearchItem(root.Left, elem)
	}
	return nil
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		if root.Left != nil {
			BTreeApplyInorder(root.Left, f)
			// f(root.Left.Data)
		}
		f(root.Data)

		if root.Right != nil {
			BTreeApplyInorder(root.Right, f)
			// f(root.Right.Data)
		}
	}
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == root {
		return rplc
	}
	current := root
	for current != nil {
		if current.Left == node {
			current.Left = rplc
			break
		}

		if current.Right == node {
			current.Right = rplc
			break
		}

		if node.Data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return root
}

func main() {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "1")
	rplc := &TreeNode{Data: "3"}
	root = BTreeTransplant(root, node, rplc)
	BTreeApplyInorder(root, fmt.Println)
}
