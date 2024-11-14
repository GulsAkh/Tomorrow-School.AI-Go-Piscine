// package piscine

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if elem == root.Data {
// 		return root
// 	}

// 	if elem < root.Data {
// 		return BTreeSearchItem(root.Left, elem)
// 	} else {
// 		return BTreeSearchItem(root.Right, elem)
// 	}
// }

// func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
// 	if root == nil {
// 		return
// 	}

// 	BTreeApplyInorder(root.Left, f)
// 	f(root.Data)
// 	BTreeApplyInorder(root.Right, f)
// }

// func findMin(node *TreeNode) *TreeNode {
// 	current := node
// 	for current.Left != nil {
// 		current = current.Left
// 	}
// 	return current
// }

// func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
// 	if node.Left == nil && node.Right == nil {
// 		if node.Parent == nil { // node is root
// 			return nil
// 		}
// 		if node.Parent.Left == node {
// 			node.Parent.Left = nil
// 		} else {
// 			node.Parent.Right = nil
// 		}
// 		return root
// 	}

// 	if node.Left == nil || node.Right == nil {
// 		var child *TreeNode
// 		if node.Left != nil {
// 			child = node.Left
// 		} else {
// 			child = node.Right
// 		}

// 		if node.Parent == nil {
// 			child.Parent = nil
// 			return child
// 		}

// 		if node.Parent.Left == node {
// 			node.Parent.Left = child
// 		} else {
// 			node.Parent.Right = child
// 		}
// 		child.Parent = node.Parent
// 		return root
// 	}

// 	successor := findMin(node.Right)
// 	node.Data = successor.Data
// 	return BTreeDeleteNode(root, successor)
// }
