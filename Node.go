package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	//"encoding/xml"
	"bufio"
	"strings"
)

type Index struct {
  data float64
  year int
}

func getRequest(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("https://data.un.org/ws/rest/data/DF_UNData_UNFCC")
  if err != nil {
    log.Fatal(err)
  }

  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  out := string(buf.Bytes())
  out = string(out[len("http://esa.un.org/unpd/wpp/index.htm"):])
  fmt.Print(out)

  
  defer resp.Body.Close()
  fmt.Fprintf(w, "%s", string(buf.Bytes()))
  reader := bufio.NewReader(strings.NewReader(out))
  for true {
    str, err := reader.ReadString('\n')
    if err != nil {
      panic(err)
    }
    str 
  }
}

func main(){
  http.HandleFunc("/view/", getRequest)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
