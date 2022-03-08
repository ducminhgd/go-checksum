package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"crypto/md5"

	"github.com/spf13/cobra"
)

const (
	bufferBytes        = 1024 * 1024 * 100 // 100MB
	CMD_CREATE  string = "create"
	CMD_VERIFY  string = "verify"
)

var (
	command, inputFile, checksum string
)

var cmdTemplate = &cobra.Command{
	Use:     "go-checksum",
	Example: "go-checksum create --input=filename.mp4",
	Run: func(cmd *cobra.Command, args []string) {
		command, _ = cmd.Flags().GetString("command")
		if command == "" {
			command = cmd.Flags().Arg(0)
		}
		inputFile, _ = cmd.Flags().GetString("input-file")
		checksum, _ = cmd.Flags().GetString("checksum")

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

		if command == CMD_VERIFY && checksum == "" {
			log.Fatalln("Missing checksum string")
		}
	},
}

func main() {
	log.SetFlags(log.LstdFlags)
	log.Println("BEGIN")

	cmdTemplate.Flags().StringP("command", "c", "", "Command: create or verify")
	cmdTemplate.Flags().StringP("input-file", "i", "", "Input file path")
	cmdTemplate.Flags().StringP("checksum", "s", "", "Checksum string to verify with input file")
	cmdTemplate.Execute()

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Cannot open file: %s, error: %v", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	digest := md5.New()

	buffer := make([]byte, bufferBytes)
	for {
		_, err := file.Read(buffer)

		// log.Printf("read %d bytes\n", bufferread)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		digest.Write(buffer)
	}

	// if _, err := io.Copy(digest, file); err != nil {
	// 	log.Fatalln(err)
	// }

	if command == CMD_CREATE {
		fmt.Println(hex.EncodeToString(digest.Sum(nil)))
		log.Println("END")
		os.Exit(0)
	}

	if command == CMD_VERIFY {
		expected := hex.EncodeToString(digest.Sum(nil))
		if expected == checksum {
			fmt.Println("OK")
		} else {
			fmt.Println("Invalid")
		}
		log.Println("END")
		os.Exit(0)
	}
	os.Exit(1)
}
