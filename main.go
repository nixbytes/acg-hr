package main

import (
  "fmt"
  "os"
  "flag"
  "strings"
  "strconv"
  "encoding/csv"
)

type User struct {
  id    int    `json:"id"`
  Name  string `json:"name"`
  Home  string `json:"home"`
  Shell string `json:"shell"`
}

func main() {
  fmt.Println("basic setup")
}


func parseFlags()(path, format string){

  flag.StringVar(&path, "path","","the path to export file")
  flag.StringVar(&format, "json","","the output format for the user information. Available optio
ns are 'csv' and 'json'. The default option is json.")

  flag.Parse()

  format = strings.ToLower(format)

  if format != "csv" && format != "json" {
    fmt.Println("Error: invalid format. Use 'json' or 'csv' instead.")
    flag.Usage()
    oss.Exit(1)
  }

  return
}
