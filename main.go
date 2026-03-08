package main

import (
	"os"
	"crypto/md5"
	"fmt"
	"cas/cas"
)

data := "Hello World 255" // Text data string

cas.Stash(data)


