package main 

import (
	fetch "github.com/FriendlyUser/texcite/pkg/fetch"
	"flag"
	"fmt"
)

func main() {
	var svar string
	flag.StringVar(&svar, "isbn", "9781451648546", "ISBN 13 code search")
	flag.Parse()
	bookString := fetch.GetBookTypes("9781451648546")
	fmt.Println(bookString)
}
