package bst

type Node struct {
	val    int
	Parent *Node
	Left   *Node
	Right  *Node
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
	Head *Node
	size int
}

func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Head: n, size: 1}
}

func (t *Tree) Insert(i int) {
	n := &Node{val: i}
	if t.Head == nil {
		t.Head = n
		t.size++
		return
	}

	h := t.Head

	for {
		if n.Compare(h) == -1 {
			if h.Left == nil {
				h.Left = n
				n.Parent = h
				break
			} else {
				h = h.Right
			}
		}
	}
	t.size++
}

func (t *Tree) Search(i int) *Node {
	h := t.Head
	n := &Node{val: i}

	for h != nil {
		switch h.Compare(n) {
		case -1:
			h = h.Right
		case 1:
			h = h.Left
		case 0:
			return h
		default:
			panic("Node not found")
		}
	}
	panic("Node not found")
}
