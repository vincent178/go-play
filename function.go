package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	htmldoc := `<html lang="en">
	<head>
			<meta charset="UTF-8">
			<title></title>
	</head>
	<body>
			
	</body>
	</html>`

	r := strings.NewReader(htmldoc)

	doc, err := html.Parse(r)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc)

	fmt.Println(doc.FirstChild)
}
