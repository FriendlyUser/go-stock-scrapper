// test cases
package main 

import (
	fetch "github.com/FriendlyUser/texcite/pkg/fetch"
	"testing"
	"fmt"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
  if a == b {
    return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestFetchDyn(t *testing.T) {
	// this spacing must match the fetch functions spacing
	expectString := `@book{book:9781451648546,
	author={Walter Isaacson},
	publisher={Simon and Schuster},
	isbn={9781451648546},
	year={2011},
	url={http://play.google.com/books/reader?id=8U2oAAAAQBAJ&hl=&printsec=frontcover&source=gbs_api}
}`
	steveJobsBook := fetch.GetBookDyn("9781451648546")
	if expectString != steveJobsBook {
		t.Logf("The string spacing is different than expected")
		t.Logf(fmt.Sprintf("%d", len(expectString)))
		t.Fatal(fmt.Sprintf("%d", len(steveJobsBook)))
	}
}

func TestFetchTypes(t *testing.T) {
	// this spacing must match the fetch functions spacing
	expectString := `@book{book:9781451648546,
	author={Walter Isaacson},
	publisher={Simon and Schuster},
	isbn={9781451648546},
	year={2011},
	url={http://play.google.com/books/reader?id=8U2oAAAAQBAJ&hl=&printsec=frontcover&source=gbs_api}
}`
	steveJobsBook := fetch.GetBookTypes("9781451648546")
	if expectString != steveJobsBook {
		fmt.Println(expectString)
		fmt.Println(steveJobsBook)
		t.Logf("The string spacing is different than expected")
		t.Logf(fmt.Sprintf("%d", len(expectString)))
		t.Fatal(fmt.Sprintf("%d", len(steveJobsBook)))
	}
}
