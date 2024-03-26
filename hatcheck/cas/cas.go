package cas

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

func stash(data string) {
	// https://go.dev/blog/strings
	data := "Hello World 255" // Text data string

	//fmt.Println(data) // prints string

	hash := md5.Sum([]byte(data)) // sum accepts []byte slice, returns []byte slice

	//fmt.Println(hash) // prints byte slice

	shardName := string(hash[0:1])
	fileName := string(hash[2:])

	fmt.Printf("Shard: %x\n", shardName)
	
	fmt.Printf("File: %x\n", fileName)
	
	objPath := "../objects/"

	// Create directory.
	path := objPath + fmt.Sprintf("%x", shardName)
	fmt.Println("Path: " + path);

	if e := os.MkdirAll(path , os.ModePerm); e != nil {
		//log.Fatal(e)
	}

	// Create file.
	//filePath := objPath + fmt.Sprintf("%x", shardName) + "/" +  fmt.Sprintf("%x", fileName)
	filePath := path + "/" +  fmt.Sprintf("%x", fileName)
	f, e := os.Create(filePath)
	check(e)
	defer f.Close()

/*
// Make file.
	fpath := path + "/file"
	f, err := Create(fpath)
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	defer f.Close()
*/
	// io.WriteString(os.Stdout, "Hello World")
	// file.WriteString
	n, e := f.WriteString(data)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

}
