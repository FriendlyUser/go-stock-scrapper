// Author David Li

package fetch 

import (
	"encoding/json"
	// "fmt"
	"net/http"
	// "io/ioutil"
	"text/template"
	"strings"
	"bytes"
)

// fetchs books from the google api with types
func GetBookTypes(isbn string) string {
	if len(isbn) == 0 {
		isbn = "9781451648546"
	}
	url := "https://www.googleapis.com/books/v1/volumes?q=isbn:" + isbn
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	gbSearch := new(GBSearch)
	json.NewDecoder(resp.Body).Decode(gbSearch)
	authors := strings.Join(gbSearch.Items[0].VolumeInfo.Authors, "")
	publisher := gbSearch.Items[0].VolumeInfo.Publisher
	publishedDate := gbSearch.Items[0].VolumeInfo.PublishDate
	readUrl := gbSearch.Items[0].AccessInfo.WebReaderLink
	bookData := BookTex{isbn, authors, publisher, publishedDate, readUrl}
	latexRef := `@book{book:{{.ISBN}},
	author={{"{"}}{{.Authors}}{{"}"}},
	publisher={{"{"}}{{.Publisher}}{{"}"}},
	isbn={{"{"}}{{.ISBN}}{{"}"}},
	year={{"{"}}{{.PublishedDate}}{{"}"}},
	url={{"{"}}{{.URL}}{{"}"}}
}`
	tmpl := template.New("tmpl")
	//parse some content and generate a template
	tmpl, err = tmpl.Parse(latexRef)
	var tmplBytes bytes.Buffer
	//merge template 'tmpl' with content of 's'
	tmpl.Execute(&tmplBytes, bookData)
	latexString := tmplBytes.String()
	return latexString
}
