package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	searchGoogle(args)
}

func searchGoogle(args []string) {
	query := strings.Join(args, "+")

	err := exec.Command("brave", fmt.Sprintf("https://google.com/search?q=%s", query)).Start()
	if err != nil {
		log.Fatal("unable to launch browser: ", err)
	}
}
