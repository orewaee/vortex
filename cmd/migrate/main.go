package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("specify action (up/down)")
		return
	}

	switch os.Args[1] {
	case "up":
		up()
		break
	case "down":
		down()
		break
	default:
		fmt.Println("invalid action. use up/down")
		break
	}
}

func up() {
	entries, err := os.ReadDir("migrations")
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !strings.Contains(info.Name(), "up") {
			continue
		}

		cmd := exec.Command("duckdb", "vortex.duckdb", "<", info.Name())
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		i += 1
	}

	fmt.Printf("complete. executed %d files\n", i)
}

func down() {
	entries, err := os.ReadDir("migrations")
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 0

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !strings.Contains(info.Name(), "down") {
			continue
		}

		cmd := exec.Command("duckdb", "vortex.duckdb", "<", info.Name())
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(stdout))
		i += 1
	}

	fmt.Printf("complete. executed %d files\n", i)
}
