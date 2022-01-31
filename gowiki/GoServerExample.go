package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Name string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Name + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}
func (p *Page) saveCsv() error {
	filename := p.Name + ".csv"
	return os.WriteFile(filename, p.Body, 0600)
}

func convToBytes(s [][]string) []byte {
	retVal := make([]byte, len(s))
	for _, element := range s {
		line := make([]byte, len(element))
		for _, at := range element {
			b := make([]byte, len(at))
			line = append(line, b...)
		}
		retVal = append(retVal, line...)
	}
	return retVal
}

func fileExists(path string) bool {
	_, err := os.Open(path)
	return err == nil
}

func loadPage(title string) (*Page, error) {
	body, err := os.ReadFile(title + ".csv")
	if err != nil {
		return nil, err
	}
	return &Page{Name: title, Body: body}, nil
}

func loadCsv(title string) (*Page, error) {
	f, err := os.Open("test.csv")
	if err != nil {
		return nil, err
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return &Page{Name: title, Body: convToBytes(lines)}, nil
}

func searchPath(path string, r *http.Request) string {
	title := r.URL.Path[len("/"+path+"/"):]
	return title
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := searchPath("view", r)
	p, err := loadPage(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Name, string(p.Body))
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := searchPath("edit", r)
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Name: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Name, p.Name, p.Body)

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := searchPath("save", r)
	body := r.FormValue("body")
	p := &Page{Name: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func csvViewHandler(w http.ResponseWriter, r *http.Request) {
	title := searchPath("view", r)
	p, err := loadCsv(title)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Name, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/viewCsv/", csvViewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
