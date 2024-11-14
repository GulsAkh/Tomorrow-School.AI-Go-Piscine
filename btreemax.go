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

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	max := BTreeMax(root)
// 	fmt.Println(max.Data)
// }
