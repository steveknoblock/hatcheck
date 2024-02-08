package main

import (
	//"io"
	//"os"
	//"log"
	"crypto/md5"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


/*
	func Create(name string) (*File, error)

	accepts string
	returns file pointer and error
*/

func main() {

	data := "Hello World"

	fmt.Println(data) // prints string

	hash := md5.Sum([]byte(data))

	fmt.Println(hash) // prints byte slice

	str := string(hash[:])
	fmt.Println(str) // prints raw string
	fmt.Printf("%x\n", str) // prints formatted string

/*
	f, e := os.Create(string(hash))
	check(e)
	defer f.Close()
*/
/*




	l, e := f.WriteString(os.Stdout, data)
	if e != nil {
		fmt.Println(e)
		return
	}
*/
}
