package main

import (
	Link "10-html-parse/Parse"
	"fmt"
	"strings"
)

var exampleHtml = `
<html>
<body>
	<div>
  		<h1>Hello!</h1>
  		<a href="/home">
    		Home
  		</a>
  		<a href="/about">About page</a>
	<div>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := Link.Parse(r)

	if err != nil {
		panic(err)
	}

	for _,link := range links {
		fmt.Println(link)
	}

}
