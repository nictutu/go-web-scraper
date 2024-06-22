package main

import (
    "fmt"
    "log"
    "net/http"
    "golang.org/x/net/html"
)

// Fetches the HTML document from the specified URL
func fetch(url string) (*html.Node, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    doc, err := html.Parse(res.Body)
    if err != nil {
        return nil, err
    }
    return doc, nil
}

// Extracts and prints all the links and their titles from the HTML document
func parseAndPrint(doc *html.Node) {
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    fmt.Println("Link:", attr.Val)
                    fmt.Println("Title:", extractText(n))
                    fmt.Println()
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
}

// Extracts the text content from an HTML node
func extractText(n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data
    }
    var text string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        text += extractText(c)
    }
    return text
}

func main() {
    url := "https://example.com"
    doc, err := fetch(url)
    if err != nil {
        log.Fatal(err)
    }
    parseAndPrint(doc)
}
