// +build ignore

package main

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-11-21

import (
	"fmt"

	"github.com/belfinor/ljatom"
)

func main() {

	for msg := range ljatom.Read() {

		fmt.Println("Time: " + msg.Created.String())
		fmt.Println("Journal: " + msg.Journal)
		fmt.Println("Post URL: " + msg.Url)
		fmt.Println("Post title: " + msg.Title)
		fmt.Print("Post body: " + msg.Content + "\n\n\n")

	}
}
