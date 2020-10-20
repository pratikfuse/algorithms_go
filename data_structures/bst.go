package data_structures
import "fmt"

type BinarySearchTree struct {
	Root *BstNode
}

type BstNode struct {
	Value int
	LeftNode *BstNode
	RightNode *BstNode
}

func GetBst() *BinarySearchTree{
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Print(value int){
	fmt.Println(value)
}

func (bst *BinarySearchTree) Traverse(node *BstNode, callback func(val int) ){
	if node == nil {
		return
	}
	bst.Traverse(node.LeftNode, bst.Print)
	callback(node.Value)
	bst.Traverse(node.RightNode, bst.Print)
}

func (bst *BinarySearchTree) Add(value int){
	if bst.Root == nil {
		node := &BstNode{
			Value:     value,
			LeftNode:  nil,
			RightNode: nil,
		}
		bst.Root = node
		return
	}

	newNode := &BstNode{
		Value:     value,
		LeftNode:  nil,
		RightNode: nil,
	}

	current := bst.Root
	var parent *BstNode

	for {
		parent = current
		if value < parent.Value {
			current = current.LeftNode

			if current == nil {
				parent.LeftNode = newNode
				break
			}
		} else {
			current = current.RightNode
			if current == nil {
				parent.RightNode = newNode
				break
			}
		}
	}
}


