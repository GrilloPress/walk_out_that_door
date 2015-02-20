package main

import (
	
	"fmt"
	"io/ioutil"
	
)

// A representation of data. Mutable -- ie you can update it or change it. The pattern. Or model of our page

type Page struct {
	
	Title string
	Body []byte
	
}

// this creates a function called save -- it takes a reciever called p and a pointer to the Page struct/instance.
// The file name is then given the name of the title p.Title and it is make into a markdown file with the extension of .md
// it finally returns by called the ioutil utility of writing a file, injects the filename and body and sets the permissions (0600)

func (p *Page) save() error {
	

	filename := p.Title + ".md"
	return ioutil.WriteFile(filename, p.Body, 0600)
	
}

// This function lets you read file. It takes the parameters of title. It then turns that into the filename with the markdow extension
// After this is takes the body and loads that through the ioutil utility
// If there are any errors it returns nil and the error message
// If all is good, it returns with the title and body and nil

func loadPage(title string) (*Page, error) {
	
	filename := title + ".md"
	body, err := ioutil.ReadFile(filename)
	
	if err != nil {
		return nil, err
	}
	
	return &Page{Title: title, Body: body}, nil
	
}

// this function which runs when the file is loaded up will create a file, save the body and title with the below values. Sve it. Then load the file via the title and print it to the screen

func main() {
	
	p1 := &Page{Title: "WorldWorstTutorial", Body: []byte("# Once upon a time... \n totally not working")}
	p1.save()
	p2, _ := loadPage("WorldWorstTutorial")
	fmt.Println(string(p2.Body))
	
}