package main

import (
	"github.com/urfave/cli"
	"fmt"
)

func actionView(c *cli.Context) error {
	solutionPaths := c.Args().Slice()

	for c, path := range solutionPaths {
		solution, err := openSolution(path)
		if err != nil {
			return cli.Exit(err, 1)
		}
	
		solution.SortParts()
	
		fmt.Println(solution)
		if c < len(solutionPaths) - 1 {
			fmt.Println() //add newline
		}
	}
	return nil
}