package piscine

// import "fmt"

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

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {

		if root.Left != nil {
			BTreeApplyPostorder(root.Left, f)
		}

		if root.Right != nil {
			BTreeApplyPostorder(root.Right, f)
		}

		f(root.Data)
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyPostorder(root, fmt.Println)
// }
