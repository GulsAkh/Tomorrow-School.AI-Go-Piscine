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

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root != nil {
		if root.Left != nil {
			if root.Left.Data < root.Data {
				BTreeIsBinary(root.Left)
				return true
			} else {
				return false
			}
		}
		if root.Right != nil {
			if root.Right.Data > root.Data {
				BTreeIsBinary(root.Right)
				return true
			} else {
				return false
			}
		}

	}
	return false
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	fmt.Println(BTreeIsBinary(root))
// }
