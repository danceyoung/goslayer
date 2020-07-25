package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danceyoung/goslayer/internal/goslayer/layer"
)

func main() {
	defer func() {
		if err := recover(); err == nil {

		}
	}()

	fmt.Println("\nGoSLayer is a tool that helps you to create a golang project in seconds.\n\n • layered base on a standard architecture layout\n • followed by Package-Oriented-Design guideline\n link: https://github.com/danceyoung/goslayer\n")

	layer := layer.NewLayer(layer.DescStep{})
	layer.JustDo("")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		layer.JustDo(scanner.Text())
	}
}
