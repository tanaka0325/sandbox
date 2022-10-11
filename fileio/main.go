package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	dirName  = "aaa"
	filename = "dummy.txt"
	content  = "dummy text"
)

func main() {
	path := fmt.Sprintf("%s/%s", dirName, filename)
	if err := ioutil.WriteFile(path, []byte(content), 0600); err != nil {
		if os.IsNotExist(err) {
			e := fmt.Errorf("need to create directory before : %w", err)
			fmt.Println(e)
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	fmt.Println("success")
}
