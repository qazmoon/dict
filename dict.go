package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	argNum := len(os.Args)
	if argNum == 1 {
		fmt.Printf("Use : dict  your_word")
		return
	}
	var word string
	for _, a := range os.Args[1:] {
		word += a
		word += " "
	}
	//word := os.Args[1]
	httpGet(word)
	fmt.Println("--------------Ok")
}

func httpGet(word string) {
	searchURL := "http://dict.youdao.com/w/" + word + "/#keyfrom=dict2.top"
	resp, err := http.Get(searchURL)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	searchResult(doc)
}

func searchResult(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "div" {

		findresult := false
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == "phrsListTab" {
				//fmt.Println(a.Val)
				findresult = true
				break
			}
		}

		if findresult {
			for node := n.FirstChild; node != nil; node = node.NextSibling {
				hasresult := false
				if node.Type == html.ElementNode && node.Data == "div" {
					for _, a := range node.Attr {
						if a.Key == "class" && a.Val == "trans-container" {
							//fmt.Println(a.Val)
							hasresult = true
							break
						}
					}
				}

				if hasresult {
					printResult(node)
					//find over
					return
				}

			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		searchResult(c)
	}
}

func printResult(n *html.Node) {
	if n == nil {
		fmt.Println("hhhhh  nil")
		return
	}

	if n.Type == html.ElementNode && n.Data == "li" {
		childnode := n.FirstChild
		fmt.Println(childnode.Data)
		printResult(n.NextSibling)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printResult(c)
	}
}
