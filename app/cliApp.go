package app

import (
	"algorithms/data_structures/bst"
	"algorithms/data_structures/linked_list"
	"fmt"

	"github.com/urfave/cli"
)

type CliApp struct {
	App *cli.App
}

func (c *CliApp) RunApp(args []string) error {

	c.App = &cli.App{
		Name: "algorithms-cli",
		Commands: []cli.Command{
			{

				Name:        "linked_list",
				Aliases:     []string{"ll"},
				Description: "run linked list",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Value: "singly",
						Usage: "Defines type of linked list",
					},
				},
				Action: func(c *cli.Context) error {
					if err, list := linked_list.GetLinkedList(c.String("type")); err != nil {
						return err
					} else {
						list.Run(c)
						return nil
					}
				},
			},
			{
				Name:        "binary-search-tree",
				Aliases:     []string{"bst"},
				Description: "run binary search tree",
				Action: func(c *cli.Context) {
					bst.GetBst().Run(c)
				},
			},
		},
	}
	fmt.Println(args, "OS Arguements")
	return c.App.Run(args)

}
