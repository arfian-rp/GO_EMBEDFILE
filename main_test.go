package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed ICON.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("icon_next.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golang nya
