package piscine

// import "fmt"

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
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	lst := []*TreeNode{root}
	i := 0
	for i < len(lst) {
		node := lst[i]
		f(node.Data)
		i++
		if node.Left != nil {
			lst = append(lst, node.Left)
		}
		if node.Right != nil {
			lst = append(lst, node.Right)
		}
	}
}

// func main() {
// 	root := &TreeNode{Data: "4"}
// 	BTreeInsertData(root, "1")
// 	BTreeInsertData(root, "7")
// 	BTreeInsertData(root, "5")
// 	BTreeApplyByLevel(root, fmt.Println)
// }
