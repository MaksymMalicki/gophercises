package main

import htmllinkparser "github.com/Maksymmalicki/gophercises/htmllinkparser/code"

func main() {
	filesPath := []string{
		"./examples/ex0.html",
		"./examples/ex1.html",
		"./examples/ex2.html",
		"./examples/ex3.html",
		"./examples/ex4.html",
	}
	for _, path := range filesPath {
		buffer := htmllinkparser.ReadExampleHTML(path)
		htmllinkparser.ParseHTML(buffer)
	}
}
