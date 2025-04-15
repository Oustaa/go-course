package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrWorkingFileNotFound = errors.New("the working file is not found")
)

func createBackup(working, backup string) error {
	_, err := os.Stat(working)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("The file you are trying to read does not exists")
			return ErrWorkingFileNotFound
		}
		return err
	}

	workFile, err := os.Open(working)
	if err != nil {
		fmt.Println("an error oucured while reding the file")
		return err
	}

	content, err := io.ReadAll(workFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func addNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(
		workingFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}

	return nil

}

func main() {
	pid := os.Getpid()
	fmt.Printf("Process ID: %d\n", pid)

	backupFile := "backupFile.txt"
	workingFile := "notes.txt"
	data := "note"

	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 1; i <= 10_000_000; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile, note)
		fmt.Println(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
