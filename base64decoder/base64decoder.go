package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage:", "base64decoder <base64EncodedString>")
	} else {
		data, err := base64.StdEncoding.DecodeString(os.Args[1])
		if err != nil {
			fmt.Println("An error was encountered decoding supplied string.")
			os.Exit(1)
		}
		fmt.Printf("%q\n", data)
	}
}
