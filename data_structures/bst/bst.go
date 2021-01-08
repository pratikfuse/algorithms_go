package bst

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type BinarySearchTree struct {
	Root *BstNode
}

type BstNode struct {
	Value     int
	LeftNode  *BstNode
	RightNode *BstNode
}

func GetBst() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Run(cli *cli.Context) {
	// create a singly linked list

	reader := bufio.NewReader(os.Stdin)
	// main loop
	for {
		fmt.Print("> ")

		byteValue, _, _ := reader.ReadLine()

		ch := strings.ToLower(string(byteValue))

		switch ch {
		case "a":
			// add to the list
			fmt.Print("Enter data: ")
			newByte, _, _ := reader.ReadLine()
			newIntValue, _ := strconv.Atoi(string(newByte))
			err := bst.Add(newIntValue)

			if err != nil {
				log.Fatal(err)
			}
		case "r":
			fmt.Print("Enter value to remove: ")
			// byteToRemove, _, _ := reader.ReadLine()
			// value, _ := strconv.Atoi(string(byteToRemove))
			// if err != nil {
			// 	log.Fatal(err)
			// }
		case "s":
			bst.Traverse(bst.Root, bst.Print)
		// case "as":
		// 	singlyList.Sort(0)
		// case "ds":
		// 	singlyList.Sort(1)
		// case "c":
		// 	singlyList.Count()
		case "x":
			fallthrough
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Error: Invalid Input!")
		}
	}

}

func (bst *BinarySearchTree) Print(value int) {
	fmt.Println(value)
}

func (bst *BinarySearchTree) Traverse(node *BstNode, callback func(val int)) {
	if node == nil {
		return
	}
	bst.Traverse(node.LeftNode, bst.Print)
	callback(node.Value)
	bst.Traverse(node.RightNode, bst.Print)
}

func (bst *BinarySearchTree) Add(value int) error {
	if bst.Root == nil {
		node := &BstNode{
			Value:     value,
			LeftNode:  nil,
			RightNode: nil,
		}
		bst.Root = node
		return nil
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
	return nil
}
