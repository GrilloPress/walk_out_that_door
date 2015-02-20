package main

import (
	
	"fmt"
	"io/ioutil"
	
)

// like a class in a way. The pattern we model our page after

type Page struct {
	
	Title string
	Body []byte
	
}

func (p *Page) save() error {
	
	filename := p.Title + ".md"
	return ioutil.WriteFile(filename, p.Body, 0600)
	
}

func loadPage(title string) (*Page, error) {
	
	filename := title + ".md"
	body, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}
	
	return &Page{Title: title, Body: body}, nil
	
}

func main() {
	
	p1 := &Page{Title: "WorldWorstTutorial", Body: []byte("# Once upon a time... \n totally not working")}
	p1.save()
	p2, _ := loadPage("WorldWorstTutorial")
	fmt.Println(string(p2.Body))
	
}