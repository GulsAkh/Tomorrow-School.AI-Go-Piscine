package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return BTreeMax(root.Right)
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	max := BTreeMax(root)
// 	fmt.Println(max.Data)
// }
