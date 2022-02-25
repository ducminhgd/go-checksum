package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
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
		log.Println(command)
		log.Println(inputFile)
		log.Println(checksum)
	},
}

func main() {
	cmdTemplate.Flags().StringP("command", "c", "", "Command: create or verify")
	cmdTemplate.Flags().StringP("input-file", "i", "", "Input file path")
	cmdTemplate.Flags().StringP("sum-string", "s", "", "Checksum string to verify with input file")
	cmdTemplate.Execute()

	os.Exit(0)
}
