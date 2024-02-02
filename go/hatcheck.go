package main

import (
  "crypto/md5"
  "fmt"
  "bufio"
  "io/ioutil"
  "os"
 )
 
func check(e error) {
  if e != nil {
      panic(e)
  }
}

func main() {
   //data = read file here
   //hash = md5.sum(data)

 

  data = "Hello World"

  hash = md5.sum(data)

  f, e := os.Create(hash)
  check(e)

  l, e := f.WriteString(data)
    if e != nil {
      fmt.Println(e)
      f.Close()
      return
  }




 }
 
 