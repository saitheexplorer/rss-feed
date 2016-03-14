package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/fatih/color"
	valid "github.com/asaskevich/govalidator"
)

type Feed struct {
	URL		string 		`valid:"requrl,required"`
}

type ParsedFeed struct {
	Title 	string 		`xml:"channel>title"`
	Posts 	[]Post 		`xml:"channel>item"`
}

type Post struct {
	Title 	string 		`xml:"title"`
	Link 	string 		`xml:"link"`
	Data	string  	`xml:"pubDate"`
}

var feed string

func (p *ParsedFeed) toString() {
	title := color.New(color.FgCyan).Add(color.Underline)

	title.Printf("%s\n\n", p.Title)

	for i, v := range p.Posts {
		fmt.Printf("%d. %s\n", i + 1, v.Title)
	}
}

func init() {
	flag.Parse()

	feed = flag.Arg(0)
}

func main() {
	f := Feed{feed}

	_, err := valid.ValidateStruct(f)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(f.URL)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	x := ParsedFeed{}

	err = xml.Unmarshal(body, &x)

	if err != nil {
		log.Fatal(err)
	}

	x.toString()
}
