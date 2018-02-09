package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func get() {
	response, _, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}
