package main

import (
	"bufio"
	"os"

	"github.com/danceyoung/goslayer/internal/goslayer/layer"
)

func main() {
	defer func() {
		if err := recover(); err == nil {

		}
	}()

	layer := layer.NewLayer(layer.DescStep{})
	layer.JustDo("")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		layer.JustDo(scanner.Text())
	}
}
