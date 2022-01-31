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
  formatLen int
}

func (p *Page) save() error {
	filename := p.Name + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}
func (p *Page) saveCsv() error {
	filename := p.Name + ".csv"
	return os.WriteFile(filename, p.Body, 0600)
}

func trimString(s string, toTrim rune) string{
  chars := []rune(s)
  for index, char := range chars {
    if char == toTrim {
      chars = append(chars[:index], chars[index+1:]...)
    }
  }
  return s
  
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
func splitAlong(c int, a []byte) [][]byte {
    r := (len(a) + c - 1) / c
    b := make([][]byte, r)
    lo, hi := 0, c
    for i := range b {
        if hi > len(a) {
            hi = len(a)
        }
        b[i] = a[lo:hi:hi]
        lo, hi = hi, hi+c
    }
    return b
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
  if err != nil {
    panic(err)
  }
  var splitSlice [][]byte = splitAlong(12, p.Body)
  
  fmt.Fprintf(w, "<h1>%s</h1>", p.Name)
	for _, element := range splitSlice {
    fmt.Fprintf(w, "<div>%s</div>", "{" + string(element) + "}")
  }
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

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
