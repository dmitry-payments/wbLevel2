package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Error changing directory:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Usage: kill <processID>")
				continue
			}
			cmd := exec.Command("kill", args[1])
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error killing process:", err)
			}
		case "ps":
			cmd := exec.Command("ps")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error running ps command:", err)
				continue
			}
			fmt.Println(string(output))
		case "exit":
			return
		default:
			cmd := exec.Command(command, args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		}
	}
}
