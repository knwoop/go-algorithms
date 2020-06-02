package bst

type Node struct {
	val                 int
	parent, left, right *Node
}

func NewNode(i int) *Node {
	return &Node{val: i}
}

func (n *Node) Compare(m *Node) int {
	if n.val < m.val {
		return -1
	} else if n.val > m.val {
		return 1
	} else {
		return 0
	}
}

type Tree struct {
	head *Node
	size int
}

func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{head: n, size: 1}
}

func (t *Tree) Insert(i int) {
	n := &Node{val: i}
	if t.head == nil {
		t.head = n
		t.size++
		return
	}

	h := t.head

	for {
		if n.Compare(h) == -1 {
			if h.left == nil {
				h.left = n
				n.parent = h
				break
			} else {
				h = h.right
			}
		}
	}
	t.size++
}

func (t *Tree) Search(i int) *Node {
	h := t.head
	n := &Node{val: i}

	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.right
		case 1:
			h = h.left
		case 0:
			return h
		default:
			panic("Node not found")
		}
	}
	panic("Node not found")
}
