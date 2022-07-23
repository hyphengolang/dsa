package evalbinarytree

// https://appliedgo.net/generictree/
// https://go.dev/blog/intro-generics
// https://pkg.go.dev/golang.org/x/exp/constraints

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T constraints.Ordered](vv []T) *TreeNode[T] {
	if len(vv) == 0 {
		return nil
	}
	root := &TreeNode[T]{Val: vv[0]}
	for i := 1; i < len(vv); i++ {
		root.Insert(vv[i])
	}
	return root
}

func (n *TreeNode[T]) Insert(val T) { // could include data and then this would take a map instead
	if n == nil {
		return
	}

	if val < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Val: val}
		} else {
			n.Left.Insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Val: val}
		} else {
			n.Right.Insert(val)
		}
	}
}

func evaluateNode[T int](root *TreeNode[T]) bool {
	if root.Left != nil {
		switch l, r := evaluateNode(root.Left), evaluateNode(root.Right); root.Val {
		case 2:
			return l || r
		default:
			return l && r
		}
	}

	return root.Val != 0
}

func TestEvaluateTreeNode(t *testing.T) {
	root := NewTreeNode([]int{2, 1, 3, 0, 1})
	if !evaluateNode(root) {
		t.Error("Expected true, got", evaluateNode(root))
	}

	root = &TreeNode[int]{Val: 0}
	if evaluateNode(root) {
		t.Error("Expected true, got", evaluateNode(root))
	}
}
