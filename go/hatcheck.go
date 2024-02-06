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

func main() {

	data := "Hello World"

	fmt.Println(data)

	hash := md5.Sum([]byte(data))

	str := string(hash[:])
	fmt.Printf("%x\n", str)
	fmt.Println(hash)
/*


	f, e := os.Create(hash)
	check(e)
	defer f.Close()

	l, e := f.WriteString(os.Stdout, data)
	if e != nil {
		fmt.Println(e)
		return
	}
*/
}
