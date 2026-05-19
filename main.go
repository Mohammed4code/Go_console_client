package main

import (

	"fmt"
)
func main()  {
	var link string
	fmt.Print("Einter URL:")
	fmt.Scan(&link)

	FetchAndDispley(link)

}
