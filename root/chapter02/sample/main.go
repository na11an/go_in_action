package main

import (
	"log"
	"os"
	// _ "github.com/goinaction/code/chapter2/sample/matchers"
	// "github.com/goinaction/code/chapter2/sample/search"
	_ "Go_study/root/chapter02/sample/matchers"
	"Go_study/root/chapter02/sample/search"
)

// execute init before main
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
