package main

import (
	//"io"
	"os"
	//"log"
	"crypto/md5"
	"fmt"
	//"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


/*
	func Create(name string) (*File, error)
	Creates a file. Will complain when given a path.
	accepts string
	returns file pointer and error
*/

func main() {
			// https://go.dev/blog/strings
	data := "Hello World" // Text data string

	//fmt.Println(data) // prints string

	hash := md5.Sum([]byte(data)) // sum accepts []byte slice, returns []byte slice

	//fmt.Println(hash) // prints byte slice

	shardName := string(hash[0:1])
	fileName := string(hash[2:])

	//fmt.Println(str) // prints raw (uninterpreted bytes https://pkg.go.dev/fmt) string
	fmt.Printf("%x\n", shardName) // prints formatted string
	fmt.Printf("%x\n", fileName)
	
	//objPath := "objects"
	//func Create(name string) (*File, error)
	//func Dir(path string) string
	//path := filepath.Dir("../objects/bar")
	// I misread this or am just stupid
	// create creates a file
	path := "../objects/"
	//f, e := os.Create(path)
	//check(e)
	//defer f.Close()

/*
	l, e := f.WriteString(os.Stdout, data)
	if e != nil {q
		fmt.Println(e)
		return
	}
*/
}
