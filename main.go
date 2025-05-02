package main

import (
	"fmt"

	"github.com/rancher/rancher/pkg/image"
)

func main() {
	fmt.Println("hello world")

	chart := image.Charts{}

	fmt.Println(chart)
}
