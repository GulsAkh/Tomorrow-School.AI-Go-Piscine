package piscine

// type TreeNode struct {
// 	Data                string
// 	Left, Right, Parent *TreeNode
// }

// func BTreeInsertData(root *TreeNode, data string) *TreeNode {
// 	if root == nil {
// 		return &TreeNode{Data: data}
// 	}
// 	if data < root.Data {
// 		root.Left = BTreeInsertData(root.Left, data)
// 	} else {
// 		root.Right = BTreeInsertData(root.Right, data)
// 	}
// 	return root
// }

func BTreeMaxHelper(root *TreeNode, parent *TreeNode) *TreeNode {
	if root == nil {
		return nil // Return nil if the tree is empty
	}

	// Traverse to the rightmost node
	if root.Right != nil {
		return BTreeMaxHelper(root.Right, root)
	}

	return root
}

// BTreeMax - Calls helper function to find the maximum node
func BTreeMax(root *TreeNode) *TreeNode {
	return BTreeMaxHelper(root, nil)
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	max := BTreeMax(root)
// 	fmt.Println(max.Data)
// }
