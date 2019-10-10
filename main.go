package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"
)

func input() (string, error) {
	var buf strings.Builder
	_, err := io.Copy(&buf, os.Stdin)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func run() error {
	// TODO
	return nil
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Print(err)
	}
}
