package main

import htmllinkparser "github.com/Maksymmalicki/gophercises/htmllinkparser/code"

func main() {
	buffer := htmllinkparser.ReadExampleHTML()
	htmllinkparser.ParseHTML(buffer)
}
