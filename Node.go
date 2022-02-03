package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	//"encoding/xml"
	//"bufio"
	//"strings"
)

type Index struct {
  data float64
  year int
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func getRequest(w http.ResponseWriter, r *http.Request) {
  client := &http.Client{}
  req, err := http.NewRequest("GET", "https://data.un.org/ws/rest/data/DF_UNData_UNFCC", nil)
  check(err)
  req.Header.Set("Accept", "text/json")
  resp, err := client.Do(req)
  check(err)  

  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  defer resp.Body.Close()
  fmt.Fprintf(w, "%s", string(buf.Bytes()))
  //reader := bufio.NewReader(strings.NewReader(out))
}

func main(){
  http.HandleFunc("/view/", getRequest)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
