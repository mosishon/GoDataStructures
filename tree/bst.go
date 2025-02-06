package tree

// Node implemented in binarytree.go
// our bst is strict bst, so we don't have to worry about duplicates
type BinarySearchTree[T any] struct {
	root *Node[T]
}

// ComparatorResult represents the result of a comparison
type ComparatorResult int

const (
	LessThan ComparatorResult = -1
	Equal    ComparatorResult = 0
	Greater  ComparatorResult = 1
)

// Type for comparator function
// because every type can have a different way of comparison
// ex: for int we can use <, >, == operators
// but for string maybe we want to compare using length or sum of ascii values or first character or ...
// so we have custom comparator function but return type is always the same
// three possible values: LessThan, Equal, Greater(same for all types)
type Comparator[T any] func(T, T) ComparatorResult

// to implement the BinaryTreeInterface we need to implement SetRoot and GetRoot methods
func (bst *BinarySearchTree[T]) SetRoot(node *Node[T]) {
	bst.root = node
}

func (bst *BinarySearchTree[T]) GetRoot() *Node[T] {
	return bst.root
}

// insert is a helper method to recursively insert a new node into the tree
func (bst *BinarySearchTree[T]) insert(current, newNode *Node[T], comparator Comparator[T]) {
	if comparator(newNode.data, current.data) < 0 { // if the data is less than the current node
		if current.Left == nil {
			current.Left = newNode
		} else {
			bst.insert(current.Left, newNode, comparator)
		}
	} else if comparator(newNode.data, current.data) == 0 { // if the data is already in the tree
		return // ignore duplicates
	} else { // if the data is greater than the current node
		if current.Right == nil {
			current.Right = newNode
		} else {
			bst.insert(current.Right, newNode, comparator)
		}
	}
}

// Insert method will insert a new node to the tree
func (bst *BinarySearchTree[T]) Insert(data T, comparator Comparator[T]) {
	newNode := &Node[T]{data: data}
	if bst.root == nil {
		bst.root = newNode
		return
	}
	bst.insert(bst.root, newNode, comparator)
}

func (bst *BinarySearchTree[T]) find(n *Node[T], data T, comparator Comparator[T]) *Node[T] {
	if n == nil {
		return nil
	}
	switch comparator(data, n.data) {
	case 0:
		return n
	case -1:
		return bst.find(n.Left, data, comparator)
	default:
		return bst.find(n.Right, data, comparator)
	}
}

func (bst *BinarySearchTree[T]) Find(data T, comparator Comparator[T]) *Node[T] {
	return bst.find(bst.root, data, comparator)
}
