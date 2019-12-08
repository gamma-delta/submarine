package main

import (
	"io/ioutil"
	"encoding/json"
	mcz "github.com/gamma-delta/molek-cheatez"
	"path/filepath"
	"errors"
)

var errBadExtension = errors.New("bad extension")

//openSolution opens a file.
//If it's a .solution file, uses molek cheatez.
//If it's a .json file, uses go's json.
func openSolution(path string) (*mcz.Solution, error) {
	solBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	if ext == ".solution" {
		return mcz.Unmarshal(solBytes)
	} else if ext == ".json" {
		out := &mcz.Solution{}
		err := json.Unmarshal(solBytes, out)
		if err != nil {
			return nil, err
		}
		return out, nil
	} else {
		return nil, errBadExtension
	}
}

//writeSolution writes a solution to a file
//Detects the kind of file to write from the extension
func writeSolution(s *mcz.Solution, path string) error {
	ext := filepath.Ext(path)
	var out []byte = nil
	if ext == ".solution" {
		out = s.UnsafeMarshal()
	} else if ext == ".json" {
		var err error
		out, err = json.MarshalIndent(s, "", "    ")
		if err != nil {
			return err
		}
	} else {
		return errBadExtension
	}

	err := ioutil.WriteFile(path, out, 0644)
	return err //if it's nil, we were just going to return nil, right?
}