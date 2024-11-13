package piscine

// type TreeNode struct {
// 	Left, Right *TreeNode
// 	Data        string
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

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	r := &TreeNode{}
	if root != nil {
		BTreeApplyInorder(r.Left, f)
		f(root.Left.Data)
		f(root.Data)

		BTreeApplyInorder(r.Right, f)
		f(root.Right.Left.Data)
		f(root.Right.Data)
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyInorder(root, fmt.Println)
// }
