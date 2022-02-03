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

func initNameQueryMap() map[string]string { 
  var m map[string]string = make(map[string]string)
  m["Greenhouse Gases"] = "DF_UNData_UNFCC"
  m["Carbon"] = m["Greenhouse Gases"] + "/A.EN_ATM_PFCE.AUs+DNK.Gg_CO2"
  return m
  
}

func getRequest(w http.ResponseWriter, r *http.Request) {
  client := &http.Client{}
  //starttime := "1995"
  //endtime := "2001"
  m := initNameQueryMap()
  req, err := http.NewRequest("GET", "https://data.un.org/ws/rest/data/" + m["Carbon"], nil)
  check(err)
  req.Header.Set("Accept", "text/json")
  resp, err := client.Do(req)
  check(err)  

  buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  defer resp.Body.Close()
  fmt.Fprintf(w, "%s", string(buf.Bytes()))
}

func main(){
  http.HandleFunc("/view/", getRequest)
  log.Fatal(http.ListenAndServe(":8080", nil))
}


