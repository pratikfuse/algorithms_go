package linked_list

import (
	"fmt"
)

type ListItem struct {
	Value int
	Next  *ListItem
}

type SinglyList struct {
	ListName string
	Head     *ListItem
}

func (list *SinglyList) Add(value int) error {
	// add a new element to the list
	listItem := &ListItem{
		Value: value,
		Next:  nil,
	}
	if list.Head == nil {
		list.Head = listItem
		return nil
	}

	currentNode := list.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = listItem
	return nil
}

func (list *SinglyList) Count() {
	if list.Head == nil {
		fmt.Println("List is empty")
		return
	}

	currentNode := list.Head
	count := 1
	for currentNode.Next != nil {
		count++
		currentNode = currentNode.Next
	}

	fmt.Printf("%d nodes in the list \n", count)
}

/**
 */
func (list *SinglyList) Remove(value int) error {
	// remove an item from list
	if list.Head == nil {
		fmt.Println("List is empty")
		return nil
	}

	temp := list.Head
	prev := temp

	// check if the data to be removed is in the head node itself
	if temp.Value == value {
		list.Head = temp.Next
		return nil
	}

	for temp != nil && temp.Value != value {
		prev = temp
		temp = temp.Next
	}

	if temp == nil {
		fmt.Printf("%d was not found in the list\n", value)
		return nil
	}

	prev.Next = temp.Next
	temp = nil
	fmt.Printf("Deleted data %d\n", value)
	return nil

}

func (list *SinglyList) FindByValue(value int) {
	// find by value
	currentNode := list.Head

	if currentNode == nil {
		fmt.Println("List is empty")
		return
	}

	position := 0
	found := false

	for currentNode.Next != nil {
		if currentNode.Value == value {
			found = true
			break
		}
		position++
	}

	if found {
		fmt.Printf("Found at position %d\n", position)
		return
	}

	fmt.Println("not found")
}

func (list *SinglyList) Show() {
	// list all the elements in the list
	currentNode := list.Head

	if currentNode == nil {
		fmt.Println("Error! The list is empty")
		return
	}
	count := 1
	for currentNode.Next != nil {
		fmt.Printf("%d: Value: %v  \n", count, currentNode.Value)
		currentNode = currentNode.Next
		count++
	}

	fmt.Printf("%d: Value: %v  \n", count, currentNode.Value)

}

func (list *SinglyList) Sort(order int) {

	if order == 0 {
		fmt.Println("Sorting in ascending order")
	} else if order == 1 {
		fmt.Println("Sorting in descending order")
	} else {
		fmt.Println("Invalid order")
		return
	}

	// sort the list
	currentNode := list.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

}
