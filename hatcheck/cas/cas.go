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

func Stash(data string) {

	hash := md5.Sum([]byte(data))

	hexHash := fmt.Sprintf("%x", hash)

	shardName := hexHash[0:2]  // first 2 hex chars (1 byte)

	fileName  := hexHash[2:]   // remaining 30 hex chars

	fmt.Printf("Shard: %s\n", shardName)
	fmt.Printf("File: %s\n", fileName)
	
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
