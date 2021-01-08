package main

import (
	"algorithms/app"
	"algorithms/data_structures/bst"
	"os"

	"github.com/urfave/cli"
)

func BinarySearchTree() {
	numbers := []int{43, 1, 411, 213, 14412, 412, 32, 41, 2134, 23, 4, 24, 24, 1, 5, 234, 5, 234, 2346, 2, 4, 53, 462, 56, 2345, 213, 451, 512, 43}
	bst := bst.GetBst()
	for _, number := range numbers {
		bst.Add(number)
	}
	bst.Traverse(bst.Root, bst.Print)
}

func main() {

	cliApp := app.CliApp{
		App: cli.NewApp(),
	}

	if err := cliApp.RunApp(os.Args); err != nil {
		os.Exit(1)
	}

}
