package main

import (
	"fmt"
	"kredit-plus/config"
)

func init() {
	config.Init()
}

func main() {
	fmt.Println("Welcome to Kredit Plus!")
}
