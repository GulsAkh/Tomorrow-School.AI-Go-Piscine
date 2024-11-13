package piscine

// import "fmt"

// type TreeNode struct {
// 	Data        string
// 	Left, Right *TreeNode
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

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		if root.Left != nil {
			BTreeApplyPreorder(root.Left, f)
		}

		if root.Right != nil {
			BTreeApplyPreorder(root.Right, f)
		}

	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyPreorder(root, fmt.Println)
// }
