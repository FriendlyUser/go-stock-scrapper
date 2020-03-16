// author David Li
// initial implementation of isbn fetching from google
// uses dynamic json parsing
// prefer javascript dynamic json parsing over golang

package pkg
import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"text/template"
	"strings"
	"bytes"
	"reflect"
)


type BookTex struct {
	ISBN          string
	Authors       string
	Publisher     string
	PublishedDate string
	URL           string
}

func ListToString(data interface{}) []string {
	if reflect.TypeOf(data).Kind() != reflect.Slice {
			panic("err: data is not slice")
	}
	slice := reflect.Indirect(reflect.ValueOf(data))
	res := make([]string, slice.Len())
	for i := 0; i < slice.Len(); i++ {
			a := slice.Index(i).Interface().(string)
			res[i] = a
	}
	return res
}

func GetBook(isbn string) string {
	if len(isbn) == 0 {
		isbn = "9781451648546"
	}
	url := "https://www.googleapis.com/books/v1/volumes?q=isbn:" + isbn
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	// dynamically reading json data
	var results map[string]interface{}
	json.Unmarshal(body, &results)
	items, ok := results["items"].([]interface{})
	if !ok {
			fmt.Println("No Items Found")
			return ""
	}

	book, ok := items[0].(map[string]interface{})
	if !ok {
		fmt.Println("Not able to find book")
		return ""
	}
	accessInfo := book["accessInfo"].(map[string]interface{})
	if accessInfo == nil {
		fmt.Println("Not able to find accessInfo")
		return ""
	}
	volumeInfo := book["volumeInfo"].(map[string]interface{})
	if volumeInfo == nil {
		fmt.Println("Not able to find accessInfo")
		return ""
	}
	fmt.Println(accessInfo)
	for key, _ := range results {
		fmt.Println("Reading Value for Key :", key)
	}
	authors := volumeInfo["authors"]
	authorsString := ""
	if authors != nil {
		authors = volumeInfo["authors"].([]interface{})
		authorsList := ListToString(authors)
		authorsString = strings.Join(authorsList," ")
	}
	publisher := volumeInfo["publisher"]
	publisherString := ""
	if publisher != nil {
		publisherString = volumeInfo["publisher"].(string)
	}
	publishedDate := volumeInfo["publishedDate"]
	dateString := ""
	if publishedDate != nil {
		dateString = volumeInfo["publishedDate"].(string)
	}
	accessString := ""
	webReaderLink := accessInfo["webReaderLink"]
	if webReaderLink != nil {
		accessString = accessInfo["webReaderLink"].(string)
	}
	// Update text template
	// https://dev.to/kirklewis/go-text-template-processing-181d
	bookData := BookTex{isbn, authorsString, publisherString, dateString, accessString}
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
