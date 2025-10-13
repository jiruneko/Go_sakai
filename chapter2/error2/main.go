package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}
}