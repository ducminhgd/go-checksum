package main

import (
	"encoding/hex"
	"log"
	"os"

	"crypto/md5"

	"github.com/spf13/cobra"
)

var (
	command, inputFile, checksum string
)

var cmdTemplate = &cobra.Command{
	Use:     "go-checksum command [OPTIONS]",
	Example: "go-checksum create --input=filename.mp4",
	Run: func(cmd *cobra.Command, args []string) {
		command, err := cmd.Flags().GetString("command")
		if err != nil {
			log.Println("Command is invalid")
		}
		inputFile, _ := cmd.Flags().GetString("input-file")
		checksum, _ := cmd.Flags().GetString("checksum")
	},
}

func main() {
	cmdTemplate.Flags().StringP("command", "c", "", "Command: create or verify")
	cmdTemplate.Flags().StringP("input-file", "i", "", "Input file path")
	cmdTemplate.Flags().StringP("sum-string", "s", "", "Checksum string to verify with input file")
	cmdTemplate.Execute()

	file, err := os.Open(inputFile)

	if err != nil {
		log.Fatalf("Cannot open file: %s, error: %v", inputFile, err)
		os.Exit(1)
	}
	defer file.Close()

	dataBuffer := make([]byte, 1024*10)
	digest := md5.New()
	for {
		digest.Write(dataBuffer)
	}
	r := digest.Sum(nil)
	return hex.EncodeToString(r)
}
