package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"

	"crypto/md5"

	"github.com/spf13/cobra"
)

var (
	command, inputFile, checksum string
	CMD_CREATE                   string = "create"
	CMD_VERIFY                   string = "verify"
)

var cmdTemplate = &cobra.Command{
	Use:     "go-checksum command [OPTIONS]",
	Example: "go-checksum create --input=filename.mp4",
	Run: func(cmd *cobra.Command, args []string) {
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			log.Println("Command is invalid")
		}
		if command == "" {
			command = args[0]
		}
		inputFile, _ := cmd.Flags().GetString("input-file")
		checksum, _ := cmd.Flags().GetString("checksum")

		cmdList := []string{CMD_CREATE, CMD_VERIFY}
		isInList := false
		for _, c := range cmdList {
			if c == command {
				isInList = true
				break
			}
		}
		if !isInList {
			log.Fatalf("Invalid command. Support: %s\n", cmdList)
		}

		if inputFile == "" {
			log.Fatalln("Missing file path")
		}
		log.Println(inputFile)

		if command == CMD_VERIFY && checksum == "" {
			log.Fatalln("Missing checksum string")
		}
	},
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	cmdTemplate.Flags().StringP("command", "c", "", "Command: create or verify")
	cmdTemplate.Flags().StringP("input-file", "i", "", "Input file path")
	cmdTemplate.Flags().StringP("sum-string", "s", "", "Checksum string to verify with input file")
	cmdTemplate.Execute()

	log.Println(inputFile)
	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("Cannot open file: %s, error: %v", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	digest := md5.New()
	if _, err := io.Copy(digest, file); err != nil {
		log.Fatalln(err)
	}
	log.Println(hex.EncodeToString(digest.Sum(nil)))
}
