package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func ipfsSave(file string) string {
	cmd := exec.Command("ipfs", "add", file)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	hash := strings.Split(string(output), " ")[1]
	return hash
}

func ipfsLoad(hash string) string {
	cmd := exec.Command("ipfs", "cat", "/ipfs/"+hash)
	output, err := cmd.CombinedOutput()
	if err != nil {
		// panic(err)
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
	}
	return string(output)
}
