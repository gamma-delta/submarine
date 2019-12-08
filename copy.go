package main

import (
	"github.com/urfave/cli"
)

func actionCopy(c *cli.Context) error {
	//Copy the first file to the second file
	path1 := c.Args().Get(0)
	path2 := c.Args().Get(1)
	if path1 == "" || path2 == "" {
		return cli.Exit("not enough filepaths given", 1)
	}

	solution, err := openSolution(path1)
	if err != nil {
		return cli.Exit(err, 1)
	}

	err = writeSolution(solution, path2)
	if err != nil {
		return cli.Exit(err, 2)
	}
	return nil
}