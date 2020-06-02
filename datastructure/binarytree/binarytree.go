package bst

type Comparable func(c1 interface{}, c2 interface{}) bool

type BinaryTree struct {
	node        interface{}
	left, right *BinaryTree

	lessFun Comparable
}

func New(compareFun Comparable) *BinaryTree {
	return &BinaryTree{
		node:    nil,
		lessFun: compareFun,
	}
}

func (tree *BinaryTree) Search(value interface{}) *BinaryTree {
	if tree.node == nil {
		return nil
	}

	if tree.node == value {
		return tree
	}
	if tree.lessFun(value, tree.node) {
		return tree.left.Search(value)
	} else {
		return tree.right.Search(value)
	}
}

func (tree *BinaryTree) Insert(value interface{}) {
	if tree.node == nil {
		tree.node = value
		tree.right = New(tree.lessFun)
		tree.left = New(tree.lessFun)
		return
	}
	if tree.lessFun(value, tree.node) {
		tree.left.Insert(value)
	} else {
		tree.right.Insert(value)
	}
}
