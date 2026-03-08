package main

import (
    "flag"
    "fmt"
)

func main() {

    /*

    hatcheck -in = "content" [string]
    hatcheck -out = "identifier" [string]

    hatcheck -save = [string]
    hatcheck -fetch = [string]

    hatcheck save [string]
    hatcheck fetch [string]

    hatcheck -stash = [string]
    hatcheck -fetch = [string]

    hat check -in = [string]
    hat check -out = [string]
    hat start

    */
/*
    numbPtr := flag.Int("numb", 42, "an int")
    forkPtr := flag.Bool("fork", false, "a bool")
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
*/
    stash := flag.String("stash", "jamsandwich", "String to save")
    fetch := flag.String("fetch", "jamsandwich", "Fetch by string")

    flag.Parse()

    fmt.Println("stash:", *stash)
    fmt.Println("fetch:", *fetch)
    //fmt.Println("fork:", *forkPtr)
    //fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}