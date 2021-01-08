package linked_list

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

type List interface {
	Add(value interface{}) error
	Remove(value interface{}) error
	Show() error
}

func ShowHelp() {
	fmt.Println("")
	fmt.Println("a  -> Add to the list")
	fmt.Println("r -> Remove from list")
	fmt.Println("s -> Show all elements")
	fmt.Println("x | q-> Exit")
	fmt.Println("**************************************")
}

func GetLinkedList(listType string) (error, *SinglyList) {
	if listType == "singly" {
		return nil, &SinglyList{
			ListName: "Singly List",
		}
	}
	return errors.New("Only supports Singly linked list"), nil
}

// command handler for cli context
func (list *SinglyList) Run(cli *cli.Context) {
	// create a singly linked list

	reader := bufio.NewReader(os.Stdin)
	// main loop
	ShowHelp()
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
			err := list.Add(newIntValue)

			if err != nil {
				log.Fatal(err)
			}
		case "r":
			fmt.Print("Enter value to remove: ")
			byteToRemove, _, _ := reader.ReadLine()
			value, _ := strconv.Atoi(string(byteToRemove))
			err := list.Remove(value)
			if err != nil {
				log.Fatal(err)
			}
		case "s":
			list.Show()
		case "as":
			list.Sort(0)
		case "ds":
			list.Sort(1)
		case "c":
			list.Count()
		case "x":
			fallthrough
		case "q":
			os.Exit(0)
		default:
			fmt.Println("Error: Invalid Input!")
			ShowHelp()
		}
	}

}
