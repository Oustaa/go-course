package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

func main() {
	timeLimit := 5 * time.Second
	fmt.Println("Press Enter to start the stopwatch...")
	_, err := fmt.Scanln() // Wait for user to press Enter
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		return
	}
	fmt.Println("Stopwatch started. Waiting for", timeLimit)
	time.Sleep(timeLimit)
	fmt.Println("Time's up! Executing the other command.")
	cmd := exec.Command("node", "test.js")
	stdoutFile, err := os.OpenFile("./stdout.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening /dev/stdout:", err)
		return
	}
	cmd.Stdout = io.MultiWriter(stdoutFile, os.Stdout)
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
	}
}
