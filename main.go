package main

import (
	//"io"
	"os"
	//"log"
	"crypto/md5"
	"fmt"
	//"path/filepath"
	"cas/cas"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

data := "Hello World 255" // Text data string

cas.Stash(data)


