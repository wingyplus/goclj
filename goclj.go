package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wingyplus/goclj/clj"
)

func main() {
	src := os.Args[1]
	b, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	f, err := clj.ParseFile(bytes.NewReader(b))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	ioutil.WriteFile(src[:len(src)-3]+".clj", []byte(f.Text()), 0644)
}
