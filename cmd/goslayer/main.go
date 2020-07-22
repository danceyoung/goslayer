package main

import (
	"bufio"
	"goslayer/internal/goslayer/layer"
	"os"
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
