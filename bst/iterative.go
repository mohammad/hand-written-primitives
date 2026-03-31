package bst

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	val   int
}

type BST struct {
	root *TreeNode
}

func (bst *BST) Insert(val int) {
	if bst.root == nil {
		bst.root = &TreeNode{
			val: val,
		}
		return
	}

	curr := bst.root
	for curr != nil {
		if curr.val > val {
			if curr.left == nil {
				curr.left = &TreeNode{
					val: val,
				}
				break
			} else {
				curr = curr.left
			}
		} else {
			if curr.right == nil {
				curr.right = &TreeNode{
					val: val,
				}
				break
			} else {
				curr = curr.right
			}
		}
	}
}

func (bst *BST) Search(val int) *TreeNode {
	if bst.root == nil {
		return nil
	}

	curr := bst.root
	for curr != nil {
		if curr.val == val {
			return curr
		}

		if curr.val > val {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	return nil
}

// Definitely did not grok this as much as I should have
func (bst *BST) InOrder() []int {
	var result []int
	stack := []*TreeNode{}

	curr := bst.root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		// get left mode
		curr = stack[len(stack)-1]
		result = append(result, curr.val)
		stack = stack[:len(stack)-1]

		curr = curr.right
	}

	return result

}
