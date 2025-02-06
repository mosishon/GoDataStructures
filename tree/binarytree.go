package tree

type BinaryTreeInterface[T any] interface {
	SetRoot(node *Node[T])
	GetRoot() *Node[T]
}

type BinaryTree[T any] struct {
	root *Node[T]
}

type Node[T any] struct {
	data  T
	Left  *Node[T]
	Right *Node[T]
}

func (bt *BinaryTree[T]) SetRoot(node *Node[T]) {
	bt.root = node
}

func (bt *BinaryTree[T]) GetRoot() *Node[T] {
	return bt.root
}

func (n *Node[T]) SetLeft(node *Node[T]) {
	n.Left = node
}
func (n *Node[T]) SetRight(node *Node[T]) {
	n.Right = node
}
func (n *Node[T]) GetData() T {
	return n.data
}
func (n *Node[T]) SetData(data T) {
	n.data = data
}

func (bt *BinaryTree[T]) Find(data T, comparator Comparator[T]) *Node[T] {
	return bt.find(bt.root, data, comparator)
}

func (bt *BinaryTree[T]) find(n *Node[T], data T, comparator Comparator[T]) *Node[T] {
	if n == nil {
		return nil
	}
	// we need comparator because we don't know the type of data for premitive types we can use == operator
	// but if we have a custom type we need to implement a comparator
	if comparator(n.GetData(), data) == Equal {
		return n
	}
	left := bt.find(n.Left, data, comparator)
	if left != nil {
		return left
	}
	return bt.find(n.Right, data, comparator)
}
