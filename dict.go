package main

import (
"net/http"
"io/ioutil"
"strings"
"fmt"
"log"
"golang.org/x/net/html"
"os"
)

func main() {
	arg_num := len(os.Args)

    if arg_num != 2 {
    	fmt.Printf("Use : dict  your_word");
    	return;
    }
    word := os.Args[1]
    httpGet(word)
    fmt.Println("--------------Ok\n")
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
    SearchResult(doc)
}


func SearchResult(n *html.Node) {
    if n.Type == html.ElementNode && n.Data == "div" {

        findresult := false;
        for _, a := range n.Attr {
            if a.Key == "id" && a.Val == "phrsListTab" {
                //fmt.Println(a.Val)

                findresult = true;
                break
            }
        }

        if findresult {
            for node := n.FirstChild; node != nil;node = node.NextSibling{

                hasresult := false
                if node.Type == html.ElementNode && node.Data == "div" {
                   for _, a := range node.Attr {
                        if a.Key == "class" && a.Val == "trans-container" {
                            //fmt.Println(a.Val)
                            hasresult = true;
                            break
                        }
                    }
                }

                if hasresult{
                    PrintResult(node)

                    //find over
                    return
                }

            }
        }
    }


    for c := n.FirstChild; c != nil; c = c.NextSibling {
        SearchResult(c)
    }
}


func PrintResult(n *html.Node){
    if n==nil{
		fmt.Println("hhhhh  nil")
        return;
    }
	

    if n.Type == html.ElementNode && n.Data == "li" {
		childnode := n.FirstChild
		fmt.Println(childnode.Data)
        PrintResult(n.NextSibling)
    }

    for c := n.FirstChild; c != nil; c = c.NextSibling {
        PrintResult(c)
    }
}
